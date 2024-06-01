constants = import_module("../../constants.star")
service_config_lib = import_module("../../lib/service_config.star")
builtins = import_module("../../lib/builtins.star")

RPC_PORT_NUM = 8545
WS_PORT_NUM = 8546
DISCOVERY_PORT_NUM = 30303
ENGINE_RPC_PORT_NUM = 8551
METRICS_PORT_NUM = 9001

# Port IDs
RPC_PORT_ID = "eth-json-rpc"
WS_PORT_ID = "eth-json-rpc-ws"
TCP_DISCOVERY_PORT_ID = "tcp-discovery"
UDP_DISCOVERY_PORT_ID = "udp-discovery"
ENGINE_RPC_PORT_ID = "engine-rpc"
ENGINE_WS_PORT_ID = "engineWs"
METRICS_PORT_ID = "metrics"



# Because structs are immutable, we pass around a map to allow full modification up until we create the final ServiceConfig
def get_default_service_config(node_struct, node_module):
    sc = service_config_lib.get_service_config_template(
        name = node_struct.el_service_name,
        image = node_struct.el_image,
        ports = node_module.USED_PORTS_TEMPLATE,
        entrypoint = node_module.ENTRYPOINT,
        cmd = node_module.CMD,
        files = node_module.FILES,
        min_cpu = DEFAULT_MIN_CPU,
        max_cpu = DEFAULT_MAX_CPU,
        min_memory = DEFAULT_MIN_MEMORY,
        max_memory = DEFAULT_MAX_MEMORY,
        labels = {
            "node_type": "execution",
        },
        node_selectors = {
            "node_preference": "fast",
        },
    )

    return sc

def upload_global_files(plan, node_modules):
    genesis_file = plan.upload_files(
        src = "../../networks/kurtosis-devnet/network-configs/genesis.json",
        name = "genesis_file",
    )

    jwt_file = plan.upload_files(
        src = constants.JWT_FILEPATH,
        name = "jwt_file",
    )

    kzg_trusted_setup_file = plan.upload_files(
        src = constants.KZG_TRUSTED_SETUP_FILEPATH,
        name = "kzg_trusted_setup",
    )

    for node_module in node_modules.values():
        for global_file in node_module.GLOBAL_FILES:
            plan.upload_files(
                src = global_file[0],
                name = global_file[1],
            )

    return jwt_file, kzg_trusted_setup_file

def get_enode_addr(plan, el_service_name):
    extract_statement = {"enode": """.result.enode | split("?") | .[0]"""}

    request_recipe = PostHttpRequestRecipe(
        endpoint = "",
        body = '{"method":"admin_nodeInfo","params":[],"id":1,"jsonrpc":"2.0"}',
        content_type = "application/json",
        port_id = RPC_PORT_ID,
        extract = extract_statement,
    )

    response = plan.request(
        service_name = el_service_name,
        recipe = request_recipe,
    )

    enode = response["extract.enode"]
    return enode

def set_max_peers(node_module, config, max_peers):
    node_module.set_max_peers(config, max_peers)
    return config

def add_bootnodes(node_module, config, bootnodes):
    if type(bootnodes) == builtins.types.list:
        if len(bootnodes) > 0:
            cmdList = config["cmd"][:]
            cmdList.append(node_module.BOOTNODE_CMD)
            config["cmd"] = cmdList

            bootnodes_str = ",".join(bootnodes)
            config["cmd"].append(bootnodes_str)
    elif type(bootnodes) == builtins.types.str:
        if len(bootnodes) > 0:
            config["cmd"].append(node_module.BOOTNODE_CMD)
            config["cmd"].append(bootnodes)
    else:
        fail("Bootnodes was not a list or string, but instead a {}", type(bootnodes))

    return config

def deploy_node(plan, config):
    service_config = service_config_lib.create_from_config(config)

    return plan.add_service(
        name = config["name"],
        config = service_config,
    )

def deploy_nodes(plan, configs):
    service_configs = {}
    for config in configs:
        service_configs[config["name"]] = service_config_lib.create_from_config(config)

    return plan.add_services(
        configs = service_configs,
    )

def generate_node_config(plan, node_modules, node_struct, bootnode_enode_addrs = []):
    node_module = node_modules[node_struct.el_type]

    # 4a. Launch EL
    el_service_config_dict = get_default_service_config(node_struct, node_module)

    if node_struct.node_type == "seed":
        el_service_config_dict = set_max_peers(node_module, el_service_config_dict, "200")
    else:
        el_service_config_dict = add_bootnodes(node_module, el_service_config_dict, bootnode_enode_addrs)

    return el_service_config_dict

def create_node(plan, node_modules, node_struct, bootnode_enode_addrs = []):
    el_service_config_dict = generate_node_config(plan, node_modules, node_struct, bootnode_enode_addrs)
    el_client_service = deploy_node(plan, el_service_config_dict)

    enode_addr = get_enode_addr(plan, node_struct.el_service_name)
    return {
        "name": node_struct.el_service_name,
        "service": el_client_service,
        "enode_addr": enode_addr,
    }

def add_metrics(metrics_enabled_services, node, el_service_name, el_client_service, node_modules):
    if node.el_type != "ethereumjs":
        metrics_enabled_services.append({
            "name": el_service_name,
            "service": el_client_service,
            "metrics_path": node_modules[node.el_type].METRICS_PATH,
        })
    return metrics_enabled_services
