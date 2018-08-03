#!/usr/bin/env ruby
# Requires msgpack-rpc: gem install msgpack-rpc
#
# To run this script, execute it from a running Nvim instance (notice the
# trailing '&' which is required since Nvim won't process events while
# running a blocking command):
#
#	:!./hello.rb &
#
# Or from another shell by setting NVIM_LISTEN_ADDRESS:
# $ NVIM_LISTEN_ADDRESS=[address] ./hello.rb

require 'msgpack/rpc'
require 'msgpack/rpc/transport/unix'
# NVIM_LISTEN_ADDRESS=127.0.0.1:6666 nvim
nvim = MessagePack::RPC::Client.new(MessagePack::RPC::UNIXTransport.new, ENV['NVIM_LISTEN_ADDRESS'])
# nvim = MessagePack::RPC::Client.new(MessagePack::RPC::UNIXTransport.new, "127.0.0.1:6666")
_, info = nvim.call('vim_get_api_info')

info["functions"].each do |f|
	next if f["name"] =~ /\Avim_/
	p "------"
	p f["name"], f
end

# p result = nvim.call(:vim_get_current_buffer)

