package com.bocabbage.learn.taco_cloud;

import lombok.Data;
import lombok.RequiredArgsConstructor;

@Data // 提供 getter/setter/equals/hashCode/toString 等
@RequiredArgsConstructor // 自动生成一个构造器，这个构造器只包含那些以 final 修饰的字段，以及那些被标记为 @NonNull 且没有初始化的字段。
public class Ingredient {
    private final String id;
    private final String name;
    private final Type type;

    public static enum Type {
        WRAP, PROTEIN, VEGGIES, CHEESE, SAUCE
    }
}
