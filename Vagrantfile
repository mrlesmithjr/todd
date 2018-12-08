# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"
Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
 
  config.vm.box = "ubuntu/trusty64"
  # config.vm.box_url = "http://cloud-images.ubuntu.com/vagrant/trusty/current/trusty-server-cloudimg-amd64-vagrant-disk1.box"
  config.vm.synced_folder '..', '/home/vagrant/go/src/github.com/toddproject', nfs: true
  
  config.vm.synced_folder '../../Mierdin', '/home/vagrant/go/src/github.com/Mierdin', type: "rsync", rsync__exclude: [".git/", "virtualenv"]

  config.vm.provider "virtualbox" do |v|
    v.memory = 2048
    v.cpus = 2
  end

  config.vm.define "todddev" do |todddev|

    todddev.vm.host_name = "todddev"
    todddev.vm.network "private_network", ip: "192.168.50.70"
    todddev.vm.network "public_network", type: "dhcp"
    todddev.vm.provision "ansible" do |ansible|
        ansible.playbook = "scripts/ansible-todd-dev.yml"
    end
  end
        
end
