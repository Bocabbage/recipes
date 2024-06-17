package com.bocabbage.newssubscriber.monodemo.web.entity;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.datatype.jsr310.deser.LocalDateTimeDeserializer;
import com.fasterxml.jackson.datatype.jsr310.ser.LocalDateTimeSerializer;
import jakarta.persistence.*;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.time.LocalDateTime;

@Data // lombok，提供 getter/setter/toString
@NoArgsConstructor
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

    @JsonSerialize(using = LocalDateTimeSerializer.class)
    @JsonDeserialize(using = LocalDateTimeDeserializer.class)
    @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd HH:mm:ss")
    private LocalDateTime createTime;

    @JsonSerialize(using = LocalDateTimeSerializer.class)
    @JsonDeserialize(using = LocalDateTimeDeserializer.class)
    @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd HH:mm:ss")
    private LocalDateTime updateTime;

    @JsonCreator
    public Article(
        @JsonProperty("id")Long id,
        @JsonProperty("uid")Long uid,
        @JsonProperty("title")String title,
        @JsonProperty("content")String content,
        @JsonProperty("createTime")LocalDateTime createTime,
        @JsonProperty("updateTime")LocalDateTime updateTime
    ) {
        this.id = id;
        this.uid = uid;
        this.title = title;
        this.content = content;
        this.createTime = createTime;
        this.updateTime = updateTime;
    }
}
