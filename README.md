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

### Steps

#### Init

Generate initial config.hcl based on charts.
not yet implemented, not curcial to main operation,
config.hcl can be writeen manually([See Sample](docs/config.hcl))

#### Plan

Compare actual releases state in cluster with desired state as described in config.hcl
Generate a plan

#### Apply

Exexcute the plan generated, install update and remove releases to reache desired state

#### Destroy
Destroy all releases
