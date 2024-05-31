package com.bocabbage.learn.taco_cloud;

import java.util.List;
import java.util.Date;

import javax.validation.constraints.NotNull;
import javax.validation.constraints.Size;

import jakarta.persistence.*;
import lombok.Data;

@Data
@Entity
public class Taco {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;

    @NotNull
    @Size(min=5, message="Name must be at least 5 characters long")
    private String name; // 产品名

    private Date createdAt;

    @ManyToMany(targetEntity = Ingredient.class)
    @Size(min=1, message="You must choose at least 1 ingredient")
    private List<String> ingredients; // 浇头

    @PrePersist // 在新实例写入数据库前执行的内容
    void onCreate() {
        createdAt = new Date();
    }
}