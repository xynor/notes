##安装python
````
brew install python3
brew install pip3 
pip3 install virtualenv 
````
##创建虚拟环境
````
notes目录下
#python3 -m venv env
#source env/bin/activate
virtualenv notepyenv --python=python3
source notepyenv/bin/activate #激活环境
deactivate #退出环境
````
##导出requirements.txt
````
pip freeze > requirements.txt
````