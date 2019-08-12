# chartadm
chartadm is a helm charts lifecycle management tool to easily deploy and maintain chart "bundles" on Kubernetes clusters in on-premises, even air-gapped environments

_This project is in early alpha stage_

Meant to be used alone or as a part of [Klusterkit](https://platform9.com/open-source/klusterkit/):


<pre>
										 
                                         
                                          /-------------------\
                                  _______ |      etcdadm      |
                                 |        \-------------------/
 /-------------------\           |        /-------------------\
 |       cctl        |   - - - - |------- |      nodeadm      |                                         
 \-------------------/           |        \-------------------/                     
                                 |_______ /-------------------\
                                          |      chartadm     |
                                          \-------------------/ 


</pre>														   


## Getting Started

### Building

```
go get -u github.com/sysbind/chartadm
```

### Plan
