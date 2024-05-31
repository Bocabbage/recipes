package com.bocabbage.learn.taco_cloud;

import java.util.List;
import java.util.Date;

import jakarta.persistence.*;
import lombok.Data;

import javax.validation.constraints.NotBlank;
import javax.validation.constraints.Digits;
import javax.validation.constraints.Pattern;
import org.hibernate.validator.constraints.CreditCardNumber;

import java.util.ArrayList;

@Data
@Entity
@Table(name = "Taco_Order")
public class Order {
    // private static final long serialVersionUID = 1L;

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;

    private Date placedAt;

    // validator 必须在控制器层使用 @Valid 来生效
    // message 为不符合要求时显示给用户的消息
    @NotBlank(message="Name is required")
    private String name;

    @NotBlank(message="Street is required")
    private String street;

    @NotBlank(message="City is required")
    private String city;

    @NotBlank(message="State is required")
    private String state;

    @NotBlank(message="Zip code is required")
    private String zip;

    @CreditCardNumber(message="Not a valid credit card number")
    private String ccNumber;

    @Pattern(regexp="^(0[1-9]|1[0-2])([\\/])([1-9][0-9])$",
            message="Must be formatted MM/YY")
    private String ccExpiration;

    @Digits(integer=3, fraction=0, message="Invalid CVV")
    private String ccCVV;

    @ManyToMany(targetEntity = Taco.class)
    private List<Taco> tacos = new ArrayList<>();

    public void addDesign(Taco design) {
        tacos.add(design);
    }

    @PrePersist
    void placeAt() {
        placedAt = new Date();
    }
}
