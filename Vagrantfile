Vagrant.configure(2) do |config|
  config.vm.box_url = "https://github.com/CommanderK5/packer-centos-template/releases/download/0.6.7/vagrant-centos-6.7.box"
  config.vm.synced_folder "app/", "/app", owner: "app", group: "app"
  config.vm.define "ws-server" do |node|
        node.vm.box = "centos6.7"
        node.vm.hostname = "ws-server"
        node.vm.network :private_network, ip: "192.168.11.100"
        node.vm.network :forwarded_port, id: "ssh", guest: 22, host: 2233
        node.vm.network :forwarded_port, guest: 80, host: 8088
  end
end
