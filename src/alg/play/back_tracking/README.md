## 回溯算法


回溯法(backtracking)是优先搜索的一种特殊情况,又称为试探法,常用于需要记录节点状
态的深度优先搜索。通常来说,排列、组合、选择类问题使用回溯法比较方便。
顾名思义,回溯法的核心是回溯。在搜索到某一节点的时候,如果我们发现目前的节点(及
其子节点)并不是需求目标时,我们回退到原来的节点继续搜索,并且把在目前节点修改的状态
还原。这样的好处是我们可以始终只对图的总状态进行修改,而非每次遍历时新建一个图来储存
状态。在具体的写法上,它与普通的深度优先搜索一样,都有 [修改当前节点状态]→[递归子节
点] 的步骤,只是多了回溯的步骤,变成了 [修改当前节点状态]→[递归子节点]→[回改当前节点
状态]。
没有接触过回溯法的读者可能会不明白我在讲什么,这也完全正常,希望以下几道题可以让
您理解回溯法。如果还是不明白,可以记住两个小诀窍,

### 链接
```text

// https://zhuanlan.zhihu.com/p/112926891
```