package com.bocabbage.newssubscriber.monodemo.web.repository;

import com.bocabbage.newssubscriber.monodemo.web.entity.Article;
import org.springframework.stereotype.Repository;
import org.springframework.transaction.annotation.Transactional;

@Repository
public interface ArticleRepo extends BaseRepo<Article, Long> {
    // 代码运行时动态生成

    @Transactional
    void deleteByUid(Long uid);

    @Transactional
    Article findByUid(long uid);
}
