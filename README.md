# GoGit
## A simple Github Toolkit written in Golang
#### How to use :
```bash
usage: gogit [--version] [--help] <command> [-u <user>] [-c <count>]
These are common gogit commands used for various sources :

to get latest events use : 
	gogit evenets -c 10                get latest public events
	gogit events -u ajangi -c 10       get latest 10 public events for user "ajangi"
````

#### Sample :
```bash
$ gogit events -u ajangi -c 30
# Output
==============================================================
======================  Github Toolkit    ====================
======================   Alireza Jangi    ====================
====================== ajangi@hotmail.com ====================
==============================================================
+----+-------------+------------------+----------+---------------------+--------------------------------------------------+
| #  |     ID      |       TYPE       | USERNAME |        REPO         |                     REPOURL                      |
+----+-------------+------------------+----------+---------------------+--------------------------------------------------+
|  0 | 14846210611 | PushEvent        | ajangi   | ajangi/gogit        | https://api.github.com/repos/ajangi/gogit        |
|  1 | 14846200437 | PushEvent        | ajangi   | ajangi/gogit        | https://api.github.com/repos/ajangi/gogit        |
|  2 | 14844131313 | PushEvent        | ajangi   | ajangi/gogit        | https://api.github.com/repos/ajangi/gogit        |
|  3 | 14844043545 | PushEvent        | ajangi   | ajangi/gogit        | https://api.github.com/repos/ajangi/gogit        |
|  4 | 14843963673 | PushEvent        | ajangi   | ajangi/gogit        | https://api.github.com/repos/ajangi/gogit        |
|  5 | 14843906882 | PushEvent        | ajangi   | ajangi/gogit        | https://api.github.com/repos/ajangi/gogit        |
|  6 | 14843738763 | PushEvent        | ajangi   | ajangi/gogit        | https://api.github.com/repos/ajangi/gogit        |
|  7 | 14843718540 | CreateEvent      | ajangi   | ajangi/gogit        | https://api.github.com/repos/ajangi/gogit        |
|  8 | 14843710922 | CreateEvent      | ajangi   | ajangi/gogit        | https://api.github.com/repos/ajangi/gogit        |
|  9 | 14793004820 | CreateEvent      | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 10 | 14793004754 | ReleaseEvent     | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 11 | 14792996889 | PushEvent        | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 12 | 14792996630 | PullRequestEvent | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 13 | 14704852096 | PushEvent        | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 14 | 14704846687 | PushEvent        | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 15 | 14704842867 | PushEvent        | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 16 | 14704839433 | PushEvent        | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 17 | 14704832974 | PushEvent        | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 18 | 14704817117 | CreateEvent      | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 19 | 14704817062 | ReleaseEvent     | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 20 | 14704810920 | PushEvent        | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 21 | 14704810697 | PullRequestEvent | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 22 | 14701519595 | DeleteEvent      | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 23 | 14701513973 | CreateEvent      | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 24 | 14701503914 | CreateEvent      | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 25 | 14701503881 | ReleaseEvent     | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 26 | 14701495467 | PushEvent        | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 27 | 14701495389 | PullRequestEvent | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 28 | 14701493798 | PullRequestEvent | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
| 29 | 14701487699 | CreateEvent      | ajangi   | SnappDoctor/drcoder | https://api.github.com/repos/SnappDoctor/drcoder |
+----+-------------+------------------+----------+---------------------+--------------------------------------------------+
```
