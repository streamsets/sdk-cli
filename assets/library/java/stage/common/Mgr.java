package com.streamsets.pipeline.stage.common.@project@;

import com.streamsets.pipeline.lib.@project@.config.*;
import com.streamsets.pipeline.@project@.tap.TapRecordHandler;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.BufferedReader;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;
import com.streamsets.pipeline.api.Stage;
import java.io.IOException;
import java.io.InputStreamReader;
import java.nio.file.Path;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.LinkedBlockingDeque;


public class @name@Mgr {
  private static final Logger LOG = LoggerFactory.getLogger(@name@Mgr.class);

  private TapImpl tapImpl;
  private Stage.Context context;
  private List<Stage.ConfigIssue> issues;
  private @name@CoreConfigBean configBean;
  private BlockingQueue<String> queue;
  private Thread producerThread;
  private ExecutorService service;
  private TapRecordHandler handler;
  private boolean shutDown = false;

  public static @name@Mgr create(Stage.Context context, @name@CoreConfigBean configBean) {
    return new @name@Mgr(context, configBean);
  }

  public @name@Mgr(Stage.Context context, @name@CoreConfigBean configBean) {
    this.context = context;
    this.configBean = configBean;
    this.tapImpl = TapImpl.create(configBean);
    this.service = Executors.newSingleThreadExecutor();
    this.handler = TapRecordHandler.newInstance(configBean.batchSize);

    queue = new LinkedBlockingDeque<String>(configBean.batchSize * 3);
  }

  public @name@Mgr verifyTap(List<Stage.ConfigIssue> issues) throws NullPointerException, IOException {
    initializeConnectionOnly(issues);
    if (!issues.isEmpty()) {
    }
    return this;
  }

  public @name@Mgr initializeConnectionOnly(List<Stage.ConfigIssue> issues) {
    this.issues = issues;
    try {
      this.tapImpl.testExecutionProcess();
    } catch (IOException e) {
      issues.add(context.createConfigIssue(
        "",
          configBean.connection.tapName,
          @name@Errors.@name@_00,
          e.getMessage()
      ));
      if (configBean.connection.configPath != null && !configBean.connection.configPath.trim().equals(""))
      {
        if (configBean.connection.exampleName == null) {
          issues.add(context.createConfigIssue(
              "name",
              "path/path",
             @name@Errors.@name@_01,
             "Error with exampleName"
          ));
        }
      }

    }


    return this;
  }

  public @name@Mgr initialize(List<Stage.ConfigIssue> issues) {
    initializeConnectionOnly(issues);
    return this;
  }


  public void startExample() {

    this.service.execute(new Runnable() {
      @Override
      public void run() {
          long max = 200;
        try {
          boolean doQuit = false;
          long expectedTimeout = calculateTimeout(configBean.timeout);
          do {
            for (int x = 0; x <= max; x++) {
                handler.add("This is a test " + x)
            }
            if (isTimeout(expectedTimeout))
              doQuit = true;

          } while (!doQuit || shutDown);
        } catch (IOException e) {
          LOG.warn(e.getMessage());
        }
      }
    });

  }


  private long calculateTimeout(int timeout) {
      return System.currentTimeMillis() + (timeout * 1000);
  }

  private boolean isTimeout(long expectedTimeout) {
    if (System.currentTimeMillis() >= expectedTimeout)
      return true;
    else
      return false;
  }
  public void stop() {
    this.shutDown = true;
    this.service.shutdown();
    while (!this.service.isShutdown()) {
      try {
        Thread.sleep(1000);
      } catch (InterruptedException e) {
        e.printStackTrace();
      }
    }
    this.handler.clear();
  }

  public String getNextRecord() {
    return this.handler.take();
//    return this.queue.poll();
  }

}
