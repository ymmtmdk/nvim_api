#!/usr/bin/env ruby

require 'msgpack/rpc'
require 'msgpack/rpc/transport/unix'
nvim = MessagePack::RPC::Client.new(MessagePack::RPC::UNIXTransport.new, ENV['NVIM_LISTEN_ADDRESS'])

_, info = nvim.call('nvim_get_api_info')

funcs = {}
info["functions"].each do |f|
	funcs[f["name"]] = f unless f["name"] =~ /\Avim_/
end

com = ARGV[0]

raise "unknown command: #{com}" unless funcs[com]

nvim.call(com) do |e|
	p e
end
