package com.bocabbage.newssubscriber.monodemo.web.controller;

import com.bocabbage.newssubscriber.monodemo.web.entity.Article;
import com.bocabbage.newssubscriber.monodemo.web.service.ArticleService;
import com.bocabbage.newssubscriber.monodemo.web.validator.ArticleParam;
import com.bocabbage.newssubscriber.monodemo.web.validator.CreateArticleValidGroup;
import com.bocabbage.newssubscriber.monodemo.web.validator.UpdateArticleValidGroup;
import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.Positive;
import jakarta.validation.constraints.PositiveOrZero;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;

import java.security.SecureRandom;

@RestController
@RequestMapping("/newssub/v1/article")
//@CrossOrigin(origins="*")
public class ArticleController {
    // For uid generation
    private static final SecureRandom secureRandom = new SecureRandom();

    @Autowired
    private ArticleService articleSvc;

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public Long createArticle(@Validated(CreateArticleValidGroup.class) @RequestBody ArticleParam articleParam) {
        var article = new Article();
        // Generate uid
        article.setUid(secureRandom.nextLong() & Long.MAX_VALUE);
        article.setTitle(articleParam.getTitle());
        article.setContent(articleParam.getContent());
        articleSvc.createArticle(article);
        return article.getUid();
    }

    // 全量资源替换，如果是部分字段替换可以用 Patch request 来实现
    @PutMapping
    public void updateArticle(@Validated(UpdateArticleValidGroup.class) @RequestBody ArticleParam articleParam) {
        var article = new Article();
        article.setUid(articleParam.getUid());
        article.setTitle(articleParam.getTitle());
        article.setContent(articleParam.getContent());
        articleSvc.updateArticle(article);
    }

    @DeleteMapping("/{uid}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public void deleteArticle(@PathVariable("uid") @Min(1) Long uid) {
        articleSvc.deleteArticle(uid);
    }

    @GetMapping("/{uid}")
    public Article getArticle(@PathVariable("uid") @Min(1) Long uid) {
        return articleSvc.getArticle(uid);
    }

    @GetMapping("/list")
    public Iterable<Article> listArticlePaging(
            @RequestParam("page") @PositiveOrZero int page,
            @RequestParam("size") @Positive int size
    ) {
        return articleSvc.listArticles(page, size).getContent();
    }

    @GetMapping("/list-all")
    public Iterable<Article> listArticleAll() {
        return articleSvc.listArticlesAll();
    }
}
