package com.bocabbage.learn.taco_cloud;

import lombok.AccessLevel;
import lombok.Data;
import lombok.RequiredArgsConstructor;
import lombok.NoArgsConstructor;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;

@Data // 提供 getter/setter/equals/hashCode/toString 等
@RequiredArgsConstructor // 自动生成一个构造器，这个构造器只包含那些以 final 修饰的字段，以及那些被标记为 @NonNull 且没有初始化的字段。
@NoArgsConstructor(access=AccessLevel.PRIVATE, force=true) // JPA 要求实体有一个无参constructor，
@Entity // JPA 实体必须添加此注解
public class Ingredient {
    @Id
    private final String id;

    private final String name;

    private final Type type;

    public static enum Type {
        WRAP, PROTEIN, VEGGIES, CHEESE, SAUCE
    }
}
