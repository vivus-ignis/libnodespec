# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"

$go_distro_url = "http://go.googlecode.com/files/go1.1.2.linux-amd64.tar.gz"

$go_install_script = <<"EOF"
  set -x
  cd /tmp
  if [ -f /etc/arch-release ] && [ ! -f /usr/bin/wget ]; then
    /bin/sh -c 'yes | pacman -Sy wget'
  fi
  [ ! -x /usr/local/go/bin/go ] && { 
    wget -qc #{$go_distro_url}
    sudo tar xf #{::File.basename $go_distro_url} -C /usr/local 
  }
  if [ -x /usr/bin/apt-get ] && [ ! -x /usr/bin/git ]; then
    apt-get install -q=2 git
  fi
  grep -q GOPATH /home/vagrant/.bashrc
  [ "$?" != 0 ] && echo 'export GOPATH=/home/vagrant/go' >> /home/vagrant/.bashrc
  mkdir -p /home/vagrant/go
  chown vagrant /home/vagrant/go
EOF

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.define "ubuntu" do |ubuntu|
    ubuntu.vm.box = "ubuntu_12_04"
    ubuntu.vm.box_url = "https://dl.dropboxusercontent.com/u/55729638/boxes/ubuntu64_12_04.box"
    ubuntu.vm.provision "shell", inline: $go_install_script
    ubuntu.vm.provider :virtualbox do |vb|
      vb.customize ["modifyvm", :id, "--memory", "256"]
    end
  end

  config.vm.define "centos" do |centos|
    centos.vm.box = "centos_6_4"
    centos.vm.box_url = "http://puppet-vagrant-boxes.puppetlabs.com/centos-64-x64-vbox4210-nocm.box"
    centos.vm.provision "shell", inline: $go_install_script
    centos.vm.provider :virtualbox do |vb|
      vb.customize ["modifyvm", :id, "--memory", "256"]
      #vb.gui = true
    end
  end

  config.vm.define "arch" do |arch|
    arch.vm.box = "archlinux_2013_07"
    arch.vm.box_url = "http://iweb.dl.sourceforge.net/project/flowboard-vagrant-boxes/arch64-2013-07-26-minimal.box"
    arch.vm.provision "shell", inline: $go_install_script
    arch.vm.provider :virtualbox do |vb|
      vb.customize ["modifyvm", :id, "--memory", "256"]
    end
  end

  config.vm.define "gentoo" do |gentoo|
    gentoo.vm.box = "gentoo_2013_06"
    gentoo.vm.box_url = "https://lxmx-vm.s3.amazonaws.com/vagrant/boxes/lxmx_gentoo-2013.05_chef-11.4.4.box"
    gentoo.vm.provision "shell", inline: $go_install_script
    gentoo.vm.provider :virtualbox do |vb|
      vb.customize ["modifyvm", :id, "--memory", "128"]
    end
  end

end
