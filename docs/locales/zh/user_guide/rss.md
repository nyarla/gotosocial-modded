# RSS

RSS 是[简易信息聚合](https://en.wikipedia.org/wiki/RSS)的缩写。这是一种在网络上共享内容的非常成熟的标准。你可能在你喜欢的新闻网站和博客上看到过这个活泼的橙色 RSS 标志：

![橙色 RSS 图标](../public/rss.svg)

如果你愿意，可以配置你的 GoToSocial 账户，以将你的贴文以 RSS 订阅源的形式发布到网上。这让没有 Fediverse 账户的人也能定期获取你的贴文更新。这非常适合使用 GoToSocial 发布长篇博客形式的贴文，并希望任何人都能轻松阅读它们的情况。

GoToSocial 账户的 RSS 源默认是关闭的。你可以通过[用户设置](./settings.md)在 `https://[你的实例域名]/settings` 启用它。

启用后，你的账户的 RSS 订阅将可在 `https://[你的实例域名]/@[你的用户名]/feed.rss` 获取。如果你使用 RSS 阅读器，可以用其打开这个地址以检查 RSS 是否正常工作。

## 哪些贴文会通过 RSS 分享？

只有你最近的 20 篇公开贴文会通过 RSS 分享。回复和转发不包括在内。不列出的贴文也不包括在内。换句话说，通过 RSS 可见的贴文将仅是通过浏览器打开你的账户页时可见的贴文。