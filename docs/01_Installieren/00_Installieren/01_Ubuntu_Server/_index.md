

## Allgemein Installieren


## Spracheinstellung  

``` 
sudo apt install locales
locale -a
sudo locale-gen de_DE.UTF-8
sudo dpkg-reconfigure locales

sudo nano /etc/locale.gen

----

de_DE.UTF-8 UTF-8

----

sudo nano /etc/default/locale

--------

LANG=de_DE.UTF-8
LANGUAGE=de_DE.UTF-8
LC_MESSAGES=de_DE.UTF-8


---------
sudo nano /etc/environment

----
LC_ALL=de_DE.UTF-8
LANG=de_DE.UTF-8

---- 
sudo update-locale
sudo locale-gen

```




```
sudo apt-get update
sudo apt-get upgrade
sudo apt-get install git-core
apt-get install  gnupg
```

## C und C++ Installieren

```

sudo apt-get install cmake gcc clang gdb build-essential git-core

```


## Mysql 
```
wget -c https://dev.mysql.com/get/mysql-apt-config_0.8.16-1_all.deb
sudo dpkg -i mysql-apt-config*
sudo apt update
sudo apt-get install mysql-server
```

## R Programmiersprache Installieren

```
sudo apt-get install r-base-dev
sudo apt-get install r-cran-evaluate


```


## PHP 
```

sudo apt-get install php php-cli php-mbstring php-xml

```
### Composer
```
apt-get install composer
```


## Python

```
sudo apt install python3 python3-dev git curl python-is-python3  python3-pip 
pip3 install ipython
pip3 install -U jupyter
jupyter
pip3 install -U  jupyterlab
jupyter-lab
pip3 install -U  notebook
jupyter notebook
pip3 install -U voila
sudo apt-get install gunicorn -y


```

## Node

* [nodejs Insdtallieren](https://github.com/nodesource/distributions/blob/master/README.md)

## nginx

### Nginx zu deinstallieren

```
sudo apt-get remove nginx nginx-common # Entfernt alle außer Konfigurationsdateien.
sudo apt-get purge nginx nginx-common # Entfernt alles.
sudo apt-get autoremove


```
