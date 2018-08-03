import mprpc
c = mprpc.RPCClient("localhost", 6666)
_, info = c.call("vim_get_api_info")
for f in info["functions"]:
    print(f["name"], f["parameters"], f["return_type"])
info = c.call("vim_get_buffers")
print(info)

