package com.tdc.demo.infrastructure.configuration;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.scheduling.annotation.EnableScheduling;

@SpringBootApplication
@ComponentScan(basePackages = "com.tdc.demo.infrastructure")
@EnableScheduling
public class TdcApplication {

	public static void main(String[] args) {
		SpringApplication.run(TdcApplication.class, args);
	}

}
