# git初始化设置

- 版本

  ```
  git version
  ```

- 获取用户名等相关信息

  ```
  git config --get user.name/user.email
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522135752142.png" alt="image-20200522135752142" style="zoom:50%;" />

- 初始化设置用户名等信息

  ```
  git config --global user.name "XXX"
  git config --global user.email "XXX"
  ```

- 解决windows和linux下git的转码问题

  ```
  git config --global core.quotepath false
  ```

- 2.23.0版本后的更新功能，让每个代码仓库使用commitGraph文件

  ```
  git config --global core.commitGraph true
  ```

- 设置每当进行非平常的垃圾清理时，git gc命令都会生成commitGraph文件

  ```
  git config --global gc.writeCommitGraph true
  ```

- ***windows***下特有设置防止报错**文件名过长**

  ```
  git config --global core.longpaths true
  ```

#### 以上均是gitea源码中git init函数处执行的命令

<hr>

# git仓库初始化

- 远程仓库连通性检查，根据返回err判断

  ```
  git ls-remote -q -h http://192.168.2.173:10080/aaron/test.git HEAD
  ```

- 初始化建立仓库（是否裸库）

  ```
  git init (--bare)
  ```

<hr>

# git维护相关

- 仓库一致性检察，会有错误报告

  ```
  git fsck
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522135431633.png" alt="image-20200522135431633" style="zoom:50%;" />

#### gitea源码中设置定时任务运行该命令，检查仓库健康状况，该命令和git gc一样比较耗时

<hr>

# git log

#### git commit的日志获取

- -N获取日志条数

```
git log -N
```

<img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522142955982.png" alt="image-20200522142955982" style="zoom:50%;" />

- --follow参数，将认定指定文件显示log，显示log包括该文件被重命名之前的log

  ```
  git log --follow fileName 
  ```

- 格式化输出log

  ```
  git log --numstat --no-merges --pretty=format:---%n%h%n%an%n%ae%n --date=iso --branches=* --first-parent
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200525164753727.png" alt="image-20200525164753727" style="zoom:50%;" />


<hr>

# git clone

参数说明

```
--bare                create a bare repository
--mirror              create a mirror repository (implies bare)
-s, --shared          setup as shared repository
--depth <depth>       create a shallow clone of that depth
-b, --branch <branch> checkout <branch> instead of the remote's HEAD
-n, --no-checkout     don't create a checkout
--quiet               Operate quietly. Progress is not reported to the standard error stream.
```

<hr>

# git pull

参数说明

```
--rebase	          本地分支与远程分支合并，并删除原本地分支，保持commit链不会出现菱形
--all	              Fetch all remotes.
```

<hr>

# git push

参数说明

```
--f	                   force push to the remote repo
```

<hr>

# git fetch

#### 获取远程仓库，本地新建dev存储

```
git fetch origin master:dev
```

源码中结合使用**git merge-base**（多路合并的最优方法；用一种三路合并的方法来查找两个提交的共同祖先；一对提交有可能有不止一个共同祖先，越晚的共同提交越好），在使用fetch之后，通过使用merge-base将两个分支最优化合并

<hr>

# git check-attr

#### （检测文件属性中是否被锁使用，LFSLocks）

参数说明

```
-a                     列出指定路径关联的所有属性
-z					   输出格式设置为机器可解析
--cache				   只考虑.gitattributes索引，忽略工作树（版本大于1.7.8）
```

<hr>

# git blame

### 追踪到具体文件/文件行的修改记录  

- 指定文件记录：

```
git blame --porcelain -- filename
```

<img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522101428406.png" alt="image-20200522101428406" style="zoom: 50%;" />

- 指定文件，不将**根提交**作为边界

  ```
  git blame --root -- fileName
  ```

- 指定文件行M-N行

  ```
  git blame -L M,N --porcelain -- fileName
  ```

对应功能：

![image-20200522102149557](C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522102149557.png)

![image-20200522102222442](C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522102222442.png)

<hr>

# commit 相关

源码中使用

```
git commit -c username -c usermail --author=username<mail>
```

- 获取最近依次更新

  ```
  git for-each-ref --sort=committerdate refs/heads/ --count 1 --format="%(committerdate)"
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522151240503.png" alt="image-20200522151240503" style="zoom:50%;" />

- 获取commit提交次数

  ```
  git rev-list --all --count -- dirpath
  git rev-list --count master
  git rev-list --count --no-merges --branches=* --date=iso --since="Mon May 18 13:44:46 2020 +0800"
  ```

<img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522104939905.png" alt="image-20200522104939905" style="zoom: 50%;" />

<img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522151459176.png" alt="image-20200522151459176" style="zoom:50%;" />

<img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200525164330828.png" alt="image-20200525164330828" style="zoom:50%;" />

- 通过前几位hash获取commit完整hash

  ```
  git rev-parse shortHash
  git rev-parse --verify hashID
  ```

<img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522110344410.png" alt="image-20200522110344410" style="zoom:50%;" />

<img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200526134256815.png" alt="image-20200526134256815" style="zoom:50%;" />

- 通过commit hash获取当前commit的详细信息

  ```
  git cat-file -p commitHash
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522113014273.png" alt="image-20200522113014273" style="zoom:50%;" />

<hr>  
# git cat-file

#### 根据hash查看代码仓库的对象信息

- 参数：

  ```
  -t hash												查看hash对象的类型
  -p hash												查看hash对象的详细信息
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200525174016749.png" alt="image-20200525174016749" style="zoom:50%;" />

<hr>

# git tag

#### Tag类似于branch,区别是branch是可以不断改变、Merge的而Tag不行。Tag可以认为是一个快照、一个记录点，用于记录某个commit点或分支的历史快照。Tag通常打开Master分支上，以保证代码的准确性。

- 基本使用：

  ```
  git tag tagName commitHash
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200525170106528.png" alt="image-20200525170106528" style="zoom:50%;" />

  查看：

  ```
  git show tagName
  ```

- 显示引用，限制为tag，将标签解除引用到对象ID中。它们将以附加的“^ {}”显示

  ```
  git show-ref --tags -d
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200525174929213.png" alt="image-20200525174929213" style="zoom:50%;" />

  或查看tag信息

  ```
  git show-ref --tags tagName
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200525175251942.png" alt="image-20200525175251942" style="zoom:50%;" />

- 参数：

  ```
  -d tagName										删除
  -a tagName -m "Message" commitHash				为标签添加额外说明
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200525170720748.png" alt="image-20200525170720748" style="zoom:50%;" />

<hr>

# git ls-file


#### 显示索引和树中文件信息

- 参数

  ```
  -z										输出不以换行结尾
  -- dirName								输出该文件夹下的文件
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200525111933321.png" alt="image-20200525111933321" style="zoom:50%;" />

<hr>

# branch相关

- 获取commit的分支名

  ```
  git name-rev commitHash
  ```

  

<img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522105616609.png" alt="image-20200522105616609" style="zoom:50%;" />

- 验证引用正确，需要输入确切路径

  ```
  git show-ref --verify -- refs/heads/master
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522155447062.png" alt="image-20200522155447062" style="zoom:50%;" />

- 显示HEAD引用指向的分支

  ```
  git symbolic-ref HEAD
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522160912534.png" alt="image-20200522160912534" style="zoom:50%;" />

- 设置默认的引用指向（HEAD指向到master）

  ```
  git symbolic-ref HEAD refs/heads/master
  ```

- git branch命令

  ```
  -D								强制删除
  -d								删除
  --contains						仅显示包含命名提交的分支
  ```

<hr>

# git remote

命令：

```
git remote rm remoteName			
git remote -v						显示所有远程名
```



<hr>

# git show

### 通过commit获取文件状态

状态分类

- A 添加
- M 修改
- D 删除

```
git show --name-status --pretty=format:"" commitHash
```

<img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522110026801.png" alt="image-20200522110026801" style="zoom:50%;" />

***获取commit的详细改动***

```
git show commithash -- fileName
```

<img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522114554720.png" alt="image-20200522114554720" style="zoom:50%;" />

<hr>

# git archive

### 创建包含指定树的树结构的指定格式的存档，并将其写入标准输出。如果指定了<前缀>，则它将预置为归档文件中的文件名。

源码中用于给客户端下载源码前的源码打包

```
git archive --prefix="file" --format=zip -o ../test1.zip commitHash
```

<img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522112357715.png" alt="image-20200522112357715" style="zoom:50%;" />

<hr>

# git diff

### 校验改动

- 两次commit之间的区别

  ```
  git diff -M commitHashA commitHashB -- filename
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522114238128.png" alt="image-20200522114238128" style="zoom:50%;" />

<img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522115147038.png" alt="image-20200522115147038" style="zoom:50%;" />

- 参数

  ```
  -p							不显示改动，只生成补丁文件
  --name-only					只显示改动文件名
  -z							异常字符出现，不使用null代替	
  --binary					生成二进制的补丁文件，该文件可被git apply使用
  ```

<hr>

# git format-patch

### 补丁

- 两次提交之间的所有patch

  ```
  git format-patch --no-signature --stdout --root commitHashA...commitHashB
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522133724467.png" alt="image-20200522133724467" style="zoom:50%;" />

- 获取截至指定commit之前所有的改动

  ```
  git format-patch --no-signature --stdout --root commitHash
  ```

<img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200522133839840.png" alt="image-20200522133839840" style="zoom:50%;" />

- 参数

  ```
  --binary							输出二进制文件，该文件可被git apply使用
  --stdout							将所有commit按照标准格式输出，而不是为每一个commit创建文件
  ```

<hr>
# git hash-object

#### 用于创建对象

- 参数：

  ```
  -w									将对象写入对象数据库（.git/objects）
  --stdin								命令行输入对象内容
  ```

- 常见使用：

  ```
  echo "sdddd" | git hash-object -w --stdin
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200525141428579.png" alt="image-20200525141428579" style="zoom:50%;" />

# git update-index

#### 索引（暂存区）操作相关

- 如果指定的文件在索引中但缺失，则将其删除。即对象在暂存区，但是对象实体已经被删除。

  ```
  git update-index --remove -z fileName
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200525114519937.png" alt="image-20200525114519937" style="zoom:50%;" />

- 将指定的对象hash添加到索引中，并以提供的文件名保存，一般结合git write-tree使用，将该文件写入树对象，之后再使用commit

  ```
  git update-index --add --replace --cacheinfo fileMode objectHash FileName
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200525120025046.png" alt="image-20200525120025046" style="zoom:50%;" />

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200525120805623.png" alt="image-20200525120805623" style="zoom:50%;" />

  **ps: 100644 表明这是一个普通文件，100755 表示一个可执行文件，120000 标识一个符号链接**

- 参数

  ```
  -z										不以换行符结束输出
  --replace								添加时覆盖同名
  --cacheinfo <mode>,<sha1>,<path>		用于注册不在当前工作目录中的文件
  ```

<hr>

# git read-tree

- git read-tree tree-ish

- 参数：

  ```
  tree-ish							是Git能够解析为树对象名称的规范，tree对象的hash，指定										tree-ish之后，会将该树对象读入索引
  --empty								不将树读入对象读入索引，而是将其清空；将暂存区清空
  ```

  ps：索引就是git add之后，文件会添加进的地方。文件为.git/index

  ​         **该命令添加tree-ish参数后仍然会将暂存区清空，不知道为什么，用法待商榷**

<hr>
# git commit-tree

- 使用示例

  ```
  
  ```

  <img src="C:\Users\linyu\Desktop\doc for git\gitDoc\src\image-20200526140020606.png" alt="image-20200526140020606" style="zoom:50%;" />

<hr>

# reference

- write-tree/read-tree相关:  [https://git-scm.com/book/zh/v2/Git-%E5%86%85%E9%83%A8%E5%8E%9F%E7%90%86-Git-%E5%AF%B9%E8%B1%A1](https://git-scm.com/book/zh/v2/Git-内部原理-Git-对象)

# 评注note和hook

