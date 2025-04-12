Vagrant.configure("2") do |config|

  config.vm.define "forge" do |forge|
    forge.vm.box = "ubuntu/trusty64"
    forge.vm.hostname = "forge"
    forge.vm.network "private_network", ip: "192.168.33.10"
  end

  config.vm.define "server" do |server|
    server.vm.box = "ubuntu/trusty64"
    server.vm.hostname = "server"
    server.vm.network "private_network", ip: "192.168.33.11"
  end
end
