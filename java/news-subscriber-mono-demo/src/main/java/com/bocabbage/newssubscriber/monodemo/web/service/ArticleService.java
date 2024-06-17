package com.bocabbage.newssubscriber.monodemo.web.service;

import com.bocabbage.newssubscriber.monodemo.web.entity.Article;
import com.bocabbage.newssubscriber.monodemo.web.repository.ArticleRepo;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cache.annotation.CacheEvict;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.cache.annotation.Caching;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class ArticleService {

    @Autowired
    private ArticleRepo articleRepo;

    public void createArticle(Article article) {
        articleRepo.save(article);
    }

    @Caching(evict = {
            @CacheEvict(value = "article", key = "#article.Uid"),
            @CacheEvict(value = "shortArticleListCache")
    })
    public void updateArticle(Article article) {
        Article oldArticle = articleRepo.findByUid(article.getUid());
        article.setId(oldArticle.getId());
        articleRepo.saveAndFlush(article);
    }

    @Caching(evict = {
            @CacheEvict(value = "article", key = "#article.Uid"),
            @CacheEvict(value = "shortArticleListCache")
    })
    public void deleteArticle(Long uid) {
        articleRepo.deleteByUid(uid);
    }

    @Cacheable(value = "article", key = "#uid", unless = "#result == null")
    public Article getArticle(Long uid) {
        return articleRepo.findByUid(uid);
    }

    public Page<Article> listArticles(int page, int size) {
        // 分页查询
        Pageable pg = PageRequest.of(page, size);
        return articleRepo.findAll(pg);
    }

    @Cacheable(value = "shortArticleListCache", unless = "#result == null")
    public List<Article> listArticlesAll() {
        return articleRepo.findAll();
    }
}
