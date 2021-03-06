sudo yum update
sudo yum install git
git config --global user.name "Jason Butcher"
git config --global user.email "jason@butcherville.com"
git clone https://github.com/jasontbutcher/hub


export RBENV_ROOT=/usr/local/rbenv
sudo mkdir $RBENV_ROOT
sudo chown $USER $RBENV_ROOT
git clone https://github.com/sstephenson/rbenv.git $RBENV_ROOT
echo 'export RBENV_ROOT=/usr/local/rbenv' >> ~/.bashrc
echo 'export PATH="$RBENV_ROOT/bin:$PATH"' >> ~/.bashrc
echo 'eval "$(rbenv init -)"' >> ~/.bashrc
[log out...]

$ ssh 18f-hub
$ cd $RBENV_ROOT
$ git pull
$ git clone https://github.com/sstephenson/ruby-build.git \
  $RBENV_ROOT/plugins/ruby-build
$ rbenv rehash
$ rbenv install -l
[pick a version...]

$ rbenv install [VERSION]
[time passes...]

$ yum groupinstall “Development Tools"

$ yum install gcc
sudo yum install -y openssl-devel readline-devel zlib-devel

$ rbenv global [VERSION]

rbenv install 2.2.3
