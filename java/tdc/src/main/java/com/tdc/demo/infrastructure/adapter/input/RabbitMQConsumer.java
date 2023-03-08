package com.tdc.demo.infrastructure.adapter.input;

import com.rabbitmq.client.Channel;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.amqp.support.AmqpHeaders;
import org.springframework.messaging.handler.annotation.Header;
import org.springframework.stereotype.Component;

import com.tdc.demo.domain.boundary.input.Consumer;
import com.tdc.demo.domain.dto.MessageDTO;

import java.io.IOException;

@Component
public class RabbitMQConsumer implements Consumer{
    private static boolean MULTIPLE = false;
    private static boolean REQUEUE = true;

    @RabbitListener(queues = "${rabbitmq.queue}", id = "listener")
    private void subscribe(MessageDTO messageDTO, Channel channel, @Header(AmqpHeaders.DELIVERY_TAG) long tag) throws IOException {
        // Simulates processing failure to send messages to the dead letter queue.
        if (tag % 10 == 0) {
            channel.basicNack(tag, MULTIPLE, REQUEUE);
            return;
        }
        this.process(messageDTO);
        channel.basicAck(tag, MULTIPLE);
    }

    public void process(MessageDTO messageDTO){
        System.out.println("Received message. Content: " +
                messageDTO.getContent() + " -- timestamp: " + messageDTO.getTimestamp()
        );
    }
}
