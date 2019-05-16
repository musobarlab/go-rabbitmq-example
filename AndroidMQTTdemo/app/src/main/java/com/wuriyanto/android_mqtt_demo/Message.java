package com.wuriyanto.android_mqtt_demo;

/**
 * Created by wurianto on 23/04/19.
 */

public class Message {

    private String from;
    private Content content;

    public Message() {}

    public Message(String from, Content content) {
        this.from = from;
        this.content = content;
    }

    public String getFrom() {
        return from;
    }

    public void setFrom(String from) {
        this.from = from;
    }

    public Content getContent() {
        return content;
    }

    public void setContent(Content content) {
        this.content = content;
    }
}
