package com.wuriyanto.android_mqtt_demo;

/**
 * Created by wurianto on 23/04/19.
 */

public class Content {

    private String header;
    private String body;

    public Content() {}

    public Content(String header, String body) {
        this.header = header;
        this.body = body;
    }

    public String getHeader() {
        return header;
    }

    public void setHeader(String header) {
        this.header = header;
    }

    public String getBody() {
        return body;
    }

    public void setBody(String body) {
        this.body = body;
    }
}
