<h1 align="center">
    <span>Labeling-system-for-RLHF</span>
</h1>

<!-- # Labeling-system-for-RLHF -->

在ChatGPT模型训练的过程中，需要标注员的介入，不断调整模型的生成内容，并再次训练，来获得更好的生成效果。为此，被拉去开发了一个标注系统的后台。顺便当作go项目的练手。
    
该项目基于`Gin`实现了一个标注系统的后台，仅需一些代码，就可以支持InstructGPT过程中的 **SFT** 以及**RLHF**的单/多轮标注。由于是内部使用，依赖服务较少，实现逻辑也更易懂。但其中为了适配一些需求且开发周期较短导致部分业务代码稍显离谱~

- 多轮标注展示
<div align="center">
<img src="/data/home/zhangyuhan/go_project/labeling-system-for-RLHF/docs/multi.png" width="700" >
</div>

可对模型生成回答进行排序与修改。排序即可转换为RLHF阶段对模型生成的回答打分。后续考虑补充部分标注标签：敏感信息（有/无）、有害性（1~5）。


# 系统介绍
- [系统演示](#系统演示)
- [基本框架](#基本框架)
- [数据的管理](#数据管理)
- [各种标注模式的支持](#标注模式的支持)
- [未来工作](#Todo)


---
## 系统演示
<div align="center">
<img src="/data/home/zhangyuhan/go_project/labeling-system-for-RLHF/docs/labeling_show.png" width="700" >
</div>


左侧为对话区，右侧为生成区
标注角色：

- A：这里代表用户，A角色的内容需要标注员手动输入

- B：代表Bot（模型），B角色内容由模型生成，标注员进行标注



## 基本框架
- 框架
<div align="center">
<img src="/data/home/zhangyuhan/go_project/labeling-system-for-RLHF/docs/label sys.png" width="700" >
</div>

- 时序图
<div align="center">
<img src="/data/home/zhangyuhan/go_project/labeling-system-for-RLHF/docs/label logic.png" width="700" >
</div>


- 数据库设计
<div align="center">
<img src="/data/home/zhangyuhan/go_project/labeling-system-for-RLHF/docs/labeled sys.png" width="700" >
</div>

   - user表：存储用户信息，绑定每次标注会话
   - datacate与dataset表：管理待标注数据集与所属数据类
   - model：管理需要基于评估的模型版本
   - sessionitem和dataitem：管理每次标注会话的信息



## 数据管理
- 待标注数据集导入格式
```
{'task':'dialog',
 'Q':['哈喽，问一下你们这是在干吗？', '诶朋友你好。', '你也看新闻呀？'], 
 'A':['我们学院组织比赛呢。', '你好你好。'],
 'TIPS': '嗯嗯，是的。'
}
```
- 标注后数据存储格式

        存储格式：二维字符串数组

        最外层括号代表整轮会话，内层表示单次会话的Q / A

        最外层顺序为对话轮次，edited_x内层顺序为标注员从优到劣的排列
```
{
 'query':[['哈喽，问一下你们这是在干吗？'], ['诶朋友你好。','噢，这个样子', '好的，再见'], ['你也看新闻呀？',...],[...],[...]], 
 'answer':[['我们学院组织比赛呢。','我们在看戏'，'没干嘛呢'], ['你好你好。',...],[...]]
 'edited_query':[['哈喽，问一下你们这是在干吗？'], ['诶朋友你好。','噢，这个样子', '好的，再见'], ['你也看新闻呀？',...],[...],[...]],
 'edited_answer':[['我们学院组织比赛呢。','我们在看戏'，'没干嘛呢'], ['你好你好。',...],[...]]，
 'user': 'xxx'
} 
```

## 支持的标注模式
- 支持单轮数据标注，可将NER, classification等任务转为prompt形式，供标注员标注。实现SFT阶段数据的标注。
- 量化标注员对模型生成的排序，可用作RLHF阶段中的reward。


## Todo
未来有机会可能会考虑如下方式的优化
- 模型推理端的batch-by-batch请求；设定每次推理只输出较短的token，没返回结束符则循环调用该接口，能给标注员及时呈现数据。
- 补充部分标注标签：敏感信息（有/无）、有害性（1~5）。
- 标注数据保证不会重复；现阶段标注数据是“有放回的抽样”形式，可以多引入一份关系表进行管理。
