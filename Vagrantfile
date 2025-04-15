Vagrant.configure("2") do |config|

  config.vm.box = "hashicorp-education/ubuntu-24-04"
  config.vm.box_version = "0.1.0"

  config.vm.define "forge" do |forge|
    forge.vm.hostname = "forge"
    forge.vm.network "private_network", ip: "192.168.33.10"
  end

  config.vm.define "server" do |server|
    server.vm.hostname = "server"
    server.vm.network "private_network", ip: "192.168.33.11"
    server.vm.network "forwarded_port", guest: 9443, host: 9443
    server.vm.network "forwarded_port", guest: 9000, host: 9000
    server.vm.synced_folder "./vagrant/server", "/home/vagrant/server"
    server.vm.provision "shell", inline: <<-SHELL
     sudo apt-get update
      sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common
      curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
      sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
      sudo apt-get update
      sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
      sudo usermod -aG docker vagrant
    SHELL
  end

  config.vm.define "deployment" do |deployment|
    deployment.vm.hostname = "deployment"
    deployment.vm.network "private_network", ip: "192.168.33.12"
    deployment.vm.provision "shell", inline: <<-SHELL
     sudo apt-get update
      sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common
      curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
      sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
      sudo apt-get update
      sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
      sudo usermod -aG docker vagrant
    SHELL
  end
end
