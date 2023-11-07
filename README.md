骑上我心爱的小摩托-Scooter
======

# 一、项目介绍

本项目是一个基于Web的短视频应用，名为Scooter，致力于为用户提供友好的短视频浏览和分析体验。
- Scooter集成了推荐算法，旨在为用户提供更优质的视频观看推荐。
- Scooter整合了语言大模型视频分析功能，不仅可以生成视频的摘要与标签，还能够进行舆论检测，为用户提供更智能的服务。
- 在前端方面，Scooter采用了Vue框架作为开发工具，实现了弹幕互动、关注和喜欢功能等操作，提供直播服务，以增加用户的参与度和互动性。
- 后端部分使用Go语言微服务框架go-zero，以Consul作为服务的注册和发现。Redis 作为缓存，MySQL进行持久化，同时使用 Elasticsearch 作为搜索引擎，以提供快速而准确的搜索功能。使用Kafka作为消息队列，实现服务之间的解耦和流量削峰。
- 系统可观测性上，Scooter引入 Jaeger 实现链路追踪，能够对请求进行细粒度的跟踪和分析。此外，使用 Prometheus 进行服务监控并接入提供Grafana面板，实现对服务性能和资源的实时监控和可视化展示。

项目详细介绍请看： https://github.com/huangsihao7/scooter-WSVA/blob/main/docs/%E6%9E%B6%E6%9E%84%E8%AE%BE%E8%AE%A1.md

# 二、功能演示

Demo: https://img.peterli.club/scooter/scooter-demo.mp4

# 三、项目分工

| 团队成员 | 主要工作 |
| -------- | -------- | 
| 王银(队长) | 后端架构设计、直播模块、用户模块、关系模块、视频流模块、推荐系统、音视频处理、代码审查|
| 黄思豪  | 评论弹幕模块、点赞模块、搜索业务、DevOps、限流、链路追踪、大语言模型视频分析、缓存设计、舆论风控 、服务监控   |
| 徐宁   |  前端开发，包括前端架构设计、前端业务模块实现等   |

# 四、项目运行说明

为了方便部署，scooter提供了Docker Compose一键部署脚本，下面是通过Docker Compose一键启动Scooter所有服务流程：
1. 安装Docker、docker-compose、显卡驱动、nvidia docker runtime等运行环境
2. 从GitHub上[clone](https://github.com/huangsihao7/scooter-WSVA)项目
3. 在项目根目录下运行`docker-compose -f docker-compose-setup.yml up`，这一步的目的是下载前后端依赖包，并编译前后端代码。方便下一步打包docker镜像。
4. 更改项目根目录下的`.env`文件，如`SPARK_APPID`,`OSS_BUCKET`,`LIVE_URL`,`ES_HOST`,`WHISPER_MODEL_PATH`**等**配置文件。确保程序能正常访问七牛云服务和星火大模型等。
5. 在项目根目录下运行`docker-compose -f docker-compose.yml up`，程序会自动下载MySQL、Redis、Gorse、Kafka、ElasticSearch、Consul、prometheus等基础环境，打包前端和后端上一步编译好的文件为Docker镜像，并启动Scooter所有服务

### 4.1 技术架构

Scooter前端使用Vue，后端使用go-zero作为微服务框架，包括API层和RPC层。API层与前端交互，提供功能中间件。RPC层实现业务逻辑，使用Consul进行服务注册和发现。存储方面，使用MySQL持久化、Redis作为缓存、Elasticsearch为搜索引擎和Kafka作为消息队列。七牛云提供视频存储和音视频分析。算法支持包括推荐算法和语言大模型。服务可观测性通过链路追踪和服务监控实现，可在Grafana展示。

<img src="docs/img/fc149451-6ce9-4a4f-b461-f0f9fe9f3e05.png" style="zoom:67%;" />

### 4.2 前端架构图

<img src="docs/img/f2f2b5b9-6a99-4325-89e3-923fd1be3025.png" style="zoom: 50%;" />


### 4.3 后端架构图

<img src="docs/img/houduanjiegou.png" style="zoom: 67%;" />

