package com.tdc.demo.infrastructure.configuration;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import org.springframework.amqp.core.*;
import org.springframework.amqp.rabbit.config.SimpleRabbitListenerContainerFactory;
import org.springframework.amqp.rabbit.connection.CachingConnectionFactory;
import org.springframework.amqp.rabbit.connection.ConnectionFactory;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.amqp.support.converter.Jackson2JsonMessageConverter;
import org.springframework.amqp.support.converter.MessageConverter;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.util.HashMap;
import java.util.Map;

@Configuration
public class BeanFactory {
    @Value("${rabbitmq.host}")
    private String host;

    @Value("${rabbitmq.virtualhost}")
    private String vhost;

    @Value("${rabbitmq.port}")
    private int port;

    @Value("${rabbitmq.username}")
    private String username;

    @Value("${rabbitmq.password}")
    private String password;

    @Value("${rabbitmq.exchange}")
    private String exchange;

    @Value("${rabbitmq.queue}")
    private String queueName;

    @Value("${spring.rabbitmq.listener.simple.prefetch}")
    private int prefetchCount;

    private static final boolean DURABLE = true;
    private static final boolean AUTO_DELETE = false;
    private static final boolean EXCLUSIVE = false;
    @Bean
    public FanoutExchange createExchange() {
        return new FanoutExchange(exchange, DURABLE, AUTO_DELETE);
    }

    @Bean
    public FanoutExchange createDLExchange() { return new FanoutExchange(exchange+"_dlx", DURABLE, AUTO_DELETE);
    }
    @Bean
    public Queue createQueue(){
        Map arguments = new HashMap<>();
        arguments.put("x-dead-letter-exchange", exchange+"_dlx");
        return new Queue(queueName, DURABLE, EXCLUSIVE, AUTO_DELETE, arguments);
    }

    @Bean
    public Queue createDLQueue() { return new Queue(queueName+"_dlq", DURABLE, EXCLUSIVE, AUTO_DELETE);}
    @Bean
    public Binding createBinding(){
        return BindingBuilder.bind(createQueue()).to(createExchange());
    }

    @Bean
    public Binding createDLBinding(){
        return BindingBuilder.bind(createDLQueue()).to(createDLExchange());
    }

    @Bean
    public MessageConverter createConverter(){
        ObjectMapper objectMapper = new ObjectMapper();
        objectMapper.registerModule(new JavaTimeModule());
        return new Jackson2JsonMessageConverter(objectMapper);
    }

    @Bean
    public ConnectionFactory createFactory(){
        CachingConnectionFactory factory = new CachingConnectionFactory();
        factory.setHost(host);
        factory.setVirtualHost(vhost);
        factory.setUsername(username);
        factory.setPassword(password);
        factory.setPort(port);
        return factory;
    }

    @Bean
    public SimpleRabbitListenerContainerFactory rabbitListenerContainerFactory(ConnectionFactory connectionFactory) {
        SimpleRabbitListenerContainerFactory factory = new SimpleRabbitListenerContainerFactory();
        factory.setConnectionFactory(connectionFactory);
        factory.setMessageConverter(createConverter());
        factory.setPrefetchCount(this.prefetchCount);
        factory.setAcknowledgeMode(AcknowledgeMode.MANUAL);
        return factory;
    }

    @Bean
    public AmqpTemplate template(ConnectionFactory connectionFactory){
        RabbitTemplate template = new RabbitTemplate(connectionFactory);
        template.setMessageConverter(createConverter());
        template.setUsePublisherConnection(true);
        return template;
    }

}
