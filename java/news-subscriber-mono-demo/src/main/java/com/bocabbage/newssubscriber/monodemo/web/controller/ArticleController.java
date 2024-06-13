package com.bocabbage.newssubscriber.monodemo.web.controller;

import com.bocabbage.newssubscriber.monodemo.web.entity.Article;
import com.bocabbage.newssubscriber.monodemo.web.service.ArticleService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.time.LocalDateTime;

@RestController
@RequestMapping("/newssub/v1/article")
//@CrossOrigin(origins="*")
public class ArticleController {

    @Autowired
    private ArticleService articleSvc;

    @PostMapping("/create")
    @ResponseStatus(HttpStatus.CREATED)
    public void createArticle(@RequestBody Article article) {
        article.setCreateTime(LocalDateTime.now());
        article.setUpdateTime(LocalDateTime.now());
        articleSvc.createArticle(article);
    }

    // 全量资源替换，如果是部分字段替换可以用 Patch request 来实现
    @PutMapping("/{uid}")
    public void updateArticle(@RequestBody Article article) {
        article.setUpdateTime(LocalDateTime.now());
        articleSvc.updateArticle(article);
    }

    @DeleteMapping("/{uid}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public void deleteArticle(@PathVariable("uid") Long uid) {
        articleSvc.deleteArticle(uid);
    }

    @GetMapping("/{uid}")
    public Article getArticle(@PathVariable("uid") Long uid) {
        return articleSvc.getArticle(uid);
    }

    @GetMapping("/list")
    public Iterable<Article> listArticlePaging(
            @RequestParam("page") int page,
            @RequestParam("size") int size
    ) {
        return articleSvc.listArticles(page, size).getContent();
    }

    @GetMapping("/list-all")
    public Iterable<Article> listArticleAll() {
        return articleSvc.listArticlesAll();
    }
}
