package com.bocabbage.newssubscriber.monodemo.web.validator;

import jakarta.validation.constraints.NotEmpty;
import jakarta.validation.constraints.NotNull;
import lombok.Data;
import org.hibernate.validator.constraints.Length;

import java.io.Serializable;

@Data
public class ArticleParam implements Serializable {

    @NotNull(message = "Empty Article UID: not allowed.", groups = {UpdateArticleValidGroup.class})
    private Long uid;

    @NotEmpty(message = "Empty Article Title: not allowed.")
    @Length(min = 1, max = 120, message = "Title length must between 1-120 character.")
    private String title;

    @NotEmpty(message = "Empty Article Title: not allowed.")
    @Length(min = 1, max = 450, message = "Content length must between 1-450 character.")
    private String content;

}
