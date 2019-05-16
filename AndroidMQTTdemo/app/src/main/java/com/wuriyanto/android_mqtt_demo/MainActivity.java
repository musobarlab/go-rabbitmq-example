package com.wuriyanto.android_mqtt_demo;

import android.app.NotificationChannel;
import android.app.NotificationManager;
import android.app.PendingIntent;
import android.content.Context;
import android.content.Intent;
import android.media.RingtoneManager;
import android.net.Uri;
import android.os.Build;
import android.support.v4.app.NotificationCompat;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;

import com.squareup.moshi.JsonAdapter;
import com.squareup.moshi.Moshi;

import org.eclipse.paho.client.mqttv3.IMqttDeliveryToken;
import org.eclipse.paho.client.mqttv3.MqttCallbackExtended;
import org.eclipse.paho.client.mqttv3.MqttException;
import org.eclipse.paho.client.mqttv3.MqttMessage;

import java.io.IOException;

public class MainActivity extends AppCompatActivity {

    private MQTTHelper mqttHelper;
    private TextView textViewMessage;
    private Button buttonSubscribe;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        mqttHelper = new MQTTHelper(this, MQTTHelper.BROKER, MQTTHelper.CLIENT_ID);
        mqttHelper.connect();

        receiveMessage();

        buttonSubscribe = (Button) findViewById(R.id.buttonSubscribe);
        textViewMessage = (TextView) findViewById(R.id.txtMessage);

        buttonSubscribe.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {

                try {
                    mqttHelper.subscribe(MQTTHelper.TOPIC, 1);
                } catch (MqttException e) {
                    Log.e(MQTTHelper.TAG, e.getMessage());
                }
            }
        });
    }

    private void receiveMessage() {
        mqttHelper.callback(new MqttCallbackExtended() {
            @Override
            public void connectComplete(boolean reconnect, String serverURI) {
                Log.i(MQTTHelper.TAG, "serverURI: "+serverURI);
            }

            @Override
            public void connectionLost(Throwable cause) {
                Log.e(MQTTHelper.TAG, "connection lost: "+cause.getMessage());
            }

            @Override
            public void messageArrived(String topic, MqttMessage mqttMessage) throws Exception {
                System.out.println("TOPIC : "+topic);

                String m = new String(mqttMessage.getPayload());

                try {
                    Message message = parseMessage(m);

                    System.out.println("From : "+ message.getFrom());
                    System.out.println("Title : "+ message.getContent().getHeader());
                    System.out.println("Body : "+ message.getContent().getBody());

                    // send notification
                    sendNotification(message.getContent().getBody(), "icon", message.getContent().getHeader());
                }catch (IOException e) {
                    Log.e(MQTTHelper.TAG, e.getMessage());
                }
            }

            @Override
            public void deliveryComplete(IMqttDeliveryToken token) {
                Log.i(MQTTHelper.TAG, "delivery complete: "+token.toString());
            }
        });
    }

    private void sendNotification(String messageBody, String icon, String title) {
        Intent intent = new Intent(this, MainActivity.class);
        intent.addFlags(Intent.FLAG_ACTIVITY_CLEAR_TOP);
        System.out.println(icon);

        PendingIntent pendingIntent = PendingIntent.getActivity(this, 0, intent, PendingIntent.FLAG_ONE_SHOT);

        String channelId = "my_channel";
        Uri defaultSoundUri = RingtoneManager.getDefaultUri(RingtoneManager.TYPE_NOTIFICATION);
        NotificationCompat.Builder notificationBuilder =
                new NotificationCompat.Builder(this, channelId)
                        .setSmallIcon(R.drawable.ic_launcher_background)
                        .setContentTitle(title)
                        .setContentText(messageBody)
                        .setAutoCancel(true)
                        .setSound(defaultSoundUri)
                        .setContentIntent(pendingIntent);

        NotificationManager notificationManager =
                (NotificationManager) getSystemService(Context.NOTIFICATION_SERVICE);

        // Since android Oreo notification channel is needed.
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            NotificationChannel channel = new NotificationChannel(channelId,
                    "Channel human readable title",
                    NotificationManager.IMPORTANCE_DEFAULT);
            notificationManager.createNotificationChannel(channel);
        }

        notificationManager.notify(0 /* ID of notification */, notificationBuilder.build());
    }

    private Message parseMessage(String json) throws IOException {
        Moshi moshi = new Moshi.Builder().build();
        JsonAdapter<Message> jsonAdapter = moshi.adapter(Message.class);

        return jsonAdapter.fromJson(json);
    }
}
