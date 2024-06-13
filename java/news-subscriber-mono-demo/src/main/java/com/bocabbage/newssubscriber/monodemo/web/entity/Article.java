package com.bocabbage.newssubscriber.monodemo.web.entity;


//import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
//import com.fasterxml.jackson.databind.annotation.JsonSerialize;
//import com.fasterxml.jackson.datatype.jsr310.deser.LocalDateTimeDeserializer;
//import com.fasterxml.jackson.datatype.jsr310.ser.LocalDateTimeSerializer;
import jakarta.persistence.*;
import lombok.Data;

import java.time.LocalDateTime;

@Data // lombok，提供 getter/setter/toString
@Entity // 将一个 Java 类标记为一个 JPA 实体类，这意味着该类的实例可以被 JPA 提供程序（如 Hibernate）持久化到数据库中
@Table(name = "article")
public class Article implements BaseEntity {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "id", nullable = false)
    private Long id;

    private Long uid;

    private String title;

    private String content;

//    @Column(name = "tags", columnDefinition = "json")
//    private ArrayList<String> tags;

//    @JsonSerialize(using = LocalDateTimeSerializer.class)
//    @JsonDeserialize(using = LocalDateTimeDeserializer.class)
    private LocalDateTime createTime;

//    @JsonSerialize(using = LocalDateTimeSerializer.class)
//    @JsonDeserialize(using = LocalDateTimeDeserializer.class)
    private LocalDateTime updateTime;
}
