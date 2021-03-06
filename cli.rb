require "neovim"
# ENV['NVIM_RUBY_LOG_LEVEL']='debug

cli = Neovim.attach_unix(ENV['NVIM_LISTEN_ADDRESS'])
buf = cli.get_current_buffer
lines = buf.lines.map{|l| l}
lines = lines.last(100).map{|l| [l, l.split(/\s/)] }.flatten.sort.uniq
lines.each do |l|
	puts l
end

