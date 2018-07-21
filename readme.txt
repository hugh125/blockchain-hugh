
远程仓库

	git remote add origin git@github.com:hugh125/blockchain-hugh.git 创建远程库
	git remote rm  origin 删除远程库
	git push -u origin master 
		第一次推送master分支时，加上了-u参数，Git不但会把本地的master分支内容推送的远程新的master分支，
		还会把本地的master分支和远程的master分支关联起来，在以后的推送或者拉取时就可以简化命令。
	git push original master 向远程库提交，提交前在本地执行(①git add;②git commit -m 'log')
	git clone  git@github.com:hugh125/gitskills.git 克隆前，远程需要存在gitskiils库

对比修改内容

	git status 查看仓库的当前状态
	git diff "file.txt"  对比当前文件与上一次修改后的差别
	要随时掌握工作区的状态，使用git status命令。
	如果git status告诉你有文件被修改过，用git diff可以查看修改内容。

撤销修改

	git checkout -- "file.txt" 把file.txt文件在工作区的修改全部撤销
	git reset HEAD "file.txt"  可以把暂存区的修改撤销掉（unstage），重新放回工作区
		场景1：当你改乱了工作区某个文件的内容，想直接丢弃工作区的修改时，用命令git checkout -- file。
		场景2：当你不但改乱了工作区某个文件的内容，还添加到了暂存区时，想丢弃修改，
			分两步，第一步用命令git reset HEAD file，就回到了场景1，第二步按场景1操作。
		场景3：已经提交了不合适的修改到版本库时，想要撤销本次提交，参考版本回退一节，不过前提是没有推送到远程库。
文件删除

	git rm "file.txt" 从版本库中删除该文件
	命令git rm用于删除一个文件。如果一个文件已经被提交到版本库，那么你永远不用担心误删，
	但是要小心，你只能恢复文件到最新版本，你会丢失最近一次提交后你修改的内容。
		
	###################################################################
	删除GiuHub上的文件夹，不删除本地
	    git rm -r --cached dirName	# --cached不会把本地的dirName删除
	    git commit -m 'delete dir dirName'
	    git push origin master
	###################################################################		
便签管理

	git tag v1.0 	创建标签
	git tag			查看所有标签
	git show v0.9	查看指定标签
	git tag -a v0.1 -m "version 0.1 released" 3628164 
		创建带有说明的标签，用-a指定标签名，-m指定说明文字
	git tag	可以查看所有标签。

	git push origin <tagname>	可以推送一个本地标签；
	git push origin --tags		可以推送全部未推送过的本地标签；
	git tag -d <tagname>		可以删除一个本地标签；
	git push origin :refs/tags/<tagname>	可以删除一个远程标签。
	
配置文件：

	.git/config	每个仓库的Git配置文件
	.gitconfig	当前用户的Git配置文件	https://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000/0013758404317281e54b6f5375640abbb11e67be4cd49e0000
	
自动上传脚本 git_push.sh

	# bash git_push.sh
	# git remote add origin git@github.com:hugh125/blockchain-hugh.git

	# 1、克隆远程仓库，修改为自己的工作路径名
	# 2、删除不需要的文件
	# 3、提交修改后的文件
	# git clone  git@github.com:hugh125/blockchain-hugh.git
	# git rm -r --cached *
	# git commit -m 'del_all_file'
	# git push origin master

	git add *
	git commit -m 'push_log'
	git push origin master
	git status
		
		