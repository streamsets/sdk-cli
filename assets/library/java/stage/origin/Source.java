package com.streamsets.pipeline.stage.origin.@project@;

import com.streamsets.pipeline.api.BatchMaker;
import com.streamsets.pipeline.config.JsonMode;
import com.streamsets.pipeline.lib.parser.DataParser;
import com.streamsets.pipeline.lib.parser.DataParserFactory;
import com.streamsets.pipeline.lib.parser.DataParserFactoryBuilder;
import com.streamsets.pipeline.lib.parser.DataParserFormat;
import com.streamsets.pipeline.api.Record;
import com.streamsets.pipeline.api.StageException;
import com.streamsets.pipeline.api.base.BaseSource;
import com.streamsets.pipeline.lib.@project@.config.@name@CoreConfigBean;

import com.streamsets.pipeline.stage.common.@project@.@name@Mgr;
import org.json.JSONArray;
import org.json.JSONObject;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.BufferedWriter;
import java.io.FileNotFoundException;
import java.io.FileWriter;
import java.io.IOException;
import java.time.Clock;
import java.time.ZonedDateTime;
import java.util.HashMap;
import java.util.Iterator;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.UUID;

public class @name@Source extends BaseSource {
  private static final Logger LOG = LoggerFactory.getLogger(@name@Source.class);
  protected final @name@CoreConfigBean configBean;
  private long totalRecords = 0;
  private @name@Mgr mgr;
  private int timeout;
  private String currentOffsetToken;
  private Map<String, SingerSchema> stream_schema;


  public @name@Source(@name@CoreConfigBean configBean) {
    this.configBean = configBean;
    mgr = new @name@Mgr(getContext(), configBean);
    this.timeout = configBean.timeout;
    this.stream_schema = new HashMap<String, SingerSchema>();
  }

  @Override
  protected List<ConfigIssue> init() {
    List<ConfigIssue> issues = super.init();
    if (issues.isEmpty()) {
      mgr.initializeConnectionOnly(issues);
    }
    this.mgr.startTap();
    return issues;
  }

  @Override
  public String produce(String lastBatchOffsetToken, int maxBatchSize, BatchMaker batchMaker) throws StageException {
    int recordsProduced = 0;
    int sdcBatchSize = calculateSdcBatchSize(maxBatchSize);
    currentOffsetToken = Objects.isNull(lastBatchOffsetToken) ? "" : lastBatchOffsetToken;

    while (recordsProduced < sdcBatchSize) {
      String record = this.mgr.getNextRecord();
      if (record != null) {
          JSONObject json = JSONObject("{\"example\":\""+record+"\"}");
            Record sdcRecord = convertToSdcRecord("example header", json);
            recordsProduced += addRecordToBatch(batchMaker, sdcRecord);
      }
    }

    return currentOffsetToken;
  }

@Override
public void destroy() {

    this.mgr.stopTap();
}

private Record convertToSdcRecord(String stream, JSONObject eRecord) {
    final String recordContext = UUID.randomUUID().toString();
    Record sdcRecord = null;
    DataParserFactoryBuilder dataParserFactoryBuilder =
        new DataParserFactoryBuilder(getContext(), DataParserFormat.JSON);
    DataParserFactory factory = dataParserFactoryBuilder
        .setMaxDataLen(1000)
        .setMode(JsonMode.MULTIPLE_OBJECTS)
        .build();
    DataParser parser = factory.getParser(recordContext, eRecord.toString().getBytes());
    try {
      sdcRecord = parser.parse();
      parser.close();
      sdcRecord.getHeader().setAttribute("My Header", stream);
    } catch(IOException e) {
      LOG.error("@name@Source::convertToSdcRecord->"+e.getMessage());
    }

    return sdcRecord;
}
  private int addRecordToBatch(BatchMaker batchMaker, Record record) {
    batchMaker.addRecord(record);
    ++totalRecords;
    return 1;
  }
  private int calculateSdcBatchSize(int maxBatchSize) {
    return Math.min(configBean.batchSize, maxBatchSize);
  }

}
