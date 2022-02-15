package com.streamsets.pipeline.stage.common.@project@;

import com.streamsets.pipeline.api.ConfigDef;
import com.streamsets.pipeline.api.ConfigDefBean;
import com.streamsets.pipeline.api.ConfigGroups;

import com.streamsets.pipeline.api.ConnectionDef;
import com.streamsets.pipeline.api.ConnectionVerifier;
import com.streamsets.pipeline.api.ConnectionVerifierDef;
import com.streamsets.pipeline.api.Dependency;
import com.streamsets.pipeline.api.HideStage;
import com.streamsets.pipeline.api.StageDef;
import com.streamsets.pipeline.api.ValueChooserModel;
import com.streamsets.pipeline.lib.@project@.config.@name@CoreConfigBean;
import com.streamsets.pipeline.lib.@project@.config.@name@Errors;
import com.streamsets.pipeline.lib.@project@.config.@name@GroupConstants;


import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

@StageDef(
    version = 1,
    label = "@name@ Connection Verifier",
    description = "Verifies @name@ Tap",
    upgraderDef = "upgrader/@name@ConnectionVerifier.yaml",
    onlineHelpRefUrl = ""
)
@HideStage(HideStage.Type.CONNECTION_VERIFIER)
@ConfigGroups(@name@ConnectionGroups.class)
@ConnectionVerifierDef(
    verifierType = @name@Connection.TYPE,
    connectionFieldName = "connection",
    connectionSelectionFieldName = "connectionSelection"
)
public class @name@ConnectionVerifier  extends ConnectionVerifier {

  @ConfigDef(
      required = true,
      type = ConfigDef.Type.MODEL,
      connectionType = @name@Connection.TYPE,
      defaultValue = ConnectionDef.Constants.CONNECTION_SELECT_MANUAL,
      label = "Connection",
      group = @name@GroupConstants.TAP,
      displayPosition = -500
  )
  @ValueChooserModel(ConnectionDef.Constants.ConnectionChooserValues.class)
  public String connectionSelection = ConnectionDef.Constants.CONNECTION_SELECT_MANUAL;

  @ConfigDefBean(
      groups = {"EXAMPLE", "ADVANCED"},
      dependencies = {
          @Dependency(
              configName = "connectionSelection",
              triggeredByValues = ConnectionDef.Constants.CONNECTION_SELECT_MANUAL
          )
      }
  )
  public @name@Connection connection;

  @Override
  protected List<ConfigIssue> initConnection() {
    List<ConfigIssue> issues = new ArrayList<>();
    @name@CoreConfigBean configBean = new @name@CoreConfigBean();
    configBean.connection = connection;
    configBean.connection.exampleName = "example connection";

    try {
      @name@Mgr.create(getContext(), configBean).verifyTap(issues);
    } catch(Exception e) {
      issues.add(getContext().createConfigIssue("","", @name@Errors.SINGERIO_00, e.getMessage()));
    }
    return issues;
  }
}
