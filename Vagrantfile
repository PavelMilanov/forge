Vagrant.configure("2") do |config|

  config.vm.box = "generic/debian12"
  config.vm.box_version = "4.3.12"

  config.vm.define "server" do |server|
    server.vm.hostname = "server"
    server.vm.network "private_network", ip: "192.168.33.11"
    server.vm.network "forwarded_port", guest: 8080, host: 8080
    server.vm.synced_folder "./vagrant/server", "/home/vagrant/server"
    server.vm.provision "shell", inline: <<-SHELL
      sudo apt inatll curl -y
      curl -fsSL https://get.docker.com -o get-docker.sh
      sudo sh ./get-docker.sh
      sudo usermod -aG docker vagrant
    SHELL
  end

  config.vm.define "deployment" do |deployment|
    deployment.vm.hostname = "deployment"
    deployment.vm.network "private_network", ip: "192.168.33.12"
    deployment.vm.provision "shell", inline: <<-SHELL
      sudo apt inatll curl -y
      curl -fsSL https://get.docker.com -o get-docker.sh
      sudo sh ./get-docker.sh
      sudo usermod -aG docker vagrant
    SHELL
  end
end
