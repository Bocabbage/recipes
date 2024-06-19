package com.bocabbage.newssubscriber.monodemo.web.validator;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;
import lombok.Data;
import org.hibernate.validator.constraints.Length;

import java.io.Serializable;

@Data
public class ArticleParam implements Serializable {

    @NotNull(message = "Param need: uid", groups = {UpdateArticleValidGroup.class})
    private Long uid;

    @NotBlank(message = "Non-empty param need: title.", groups = {CreateArticleValidGroup.class, UpdateArticleValidGroup.class})
    @Length(min = 1, max = 120, message = "Title length must between 1-120 character.")
    private String title;

    @NotBlank(message = "Non-empty param need: content.", groups = {CreateArticleValidGroup.class, UpdateArticleValidGroup.class})
    @Length(min = 1, max = 450, message = "Content length must between 1-450 character.")
    private String content;

}
