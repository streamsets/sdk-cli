package com.streamsets.pipeline.@project@.example;

import java.util.concurrent.LinkedBlockingQueue;

public class @name@RecordHandler {
  private boolean destroy = false;
  private LinkedBlockingQueue<String> records = null;
  private static @name@RecordHandler instance = null;
  private final int _BATCH_MULTIPLIER = 3;
  public static synchronized @name@RecordHandler newInstance(int batchSize) {
    if (instance == null)
      instance = new @name@RecordHandler(batchSize);

    return instance;
  }

  private @name@RecordHandler(int batchSize) {
    records = new LinkedBlockingQueue<String>(batchSize * this._BATCH_MULTIPLIER);
  }

 public void add(String record) {
    try {
      records.put(record);
    } catch(InterruptedException e) {
      System.out.println(e.getMessage());
    }
 }

 public String take() {
    try {
      return records.take();
    } catch(InterruptedException e) {
      System.out.println(e.getMessage());
    }

    return null;
 }

 public void clear() {
    synchronized (this) {
      this.records.clear();
    }
 }
 public int size() {
    synchronized (this) {
      return records.size();
    }
 }
}
