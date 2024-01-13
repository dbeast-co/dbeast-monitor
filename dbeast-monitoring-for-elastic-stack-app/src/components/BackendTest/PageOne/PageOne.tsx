import * as React from "react";
import { getBackendSrv } from "@grafana/runtime";
import { useAsync } from "react-use";
import {  HorizontalGroup } from "@grafana/ui";
import  {testIds} from "./testIds"

export const PageOne = () => {
  const { error, loading, value } = useAsync(() => {
    const backendSrv = getBackendSrv();


    const credentials  = {
      prod: {
        elasticsearch: {
          host: "https://35.170.69.248:9200",
          authentication_enabled: true,
          username: "monitoring",
          password: "qwe123",
          status: "UNKNOWN"
        },
        kibana: {
          host: "http://35.170.69.248:5601",
          authentication_enabled: true,
          username: "monitoring",
          password: "qwe123",
          status: "UNKNOWN"
        }
      },
      mon: {
        elasticsearch: {
          host: "https://35.170.69.248:9200",
          authentication_enabled: true,
          username: "monitoring",
          password: "qwe123",
          status: "UNKNOWN"
        }
      }
    };



    return Promise.all([
      // backendSrv.get(`api/plugins/dbeast-dbeastmonitor-app/resources/ping`),
      // backendSrv.get(`api/plugins/dbeast-dbeastmonitor-app/health`),
      backendSrv.post(`api/plugins/dbeast-dbeastmonitor-app/resources/test_cluster`, credentials),
        backendSrv.post(`api/plugins/dbeast-dbeastmonitor-app/resources/save`, credentials)
    ]);
  });

  if (loading) {
    return (
      <div data-testid={testIds.pageOne.container}>
        <span>Loading...</span>
      </div>
    );
  }

  if (error || !value) {
    return (
      <div data-testid={testIds.pageOne.container}>
        <span>Error: {error?.message}</span>
      </div>
    );
  }

  const [statusData, templates] = value;

  return (
    <div data-testid={testIds.pageOne.container}>
      {/*<HorizontalGroup>*/}
      {/*  <h3>Plugin Health Check</h3>{" "}*/}
      {/*  <span data-testid={testIds.pageOne.health}>*/}
      {/*    {renderHealth(health?.message)}*/}
      {/*  </span>*/}
      {/*</HorizontalGroup>*/}
      {/*<HorizontalGroup>*/}
      {/*  <h3>Ping Backend</h3>{" "}*/}
      {/*  <span data-testid={testIds.pageOne.ping }>{ping?.message}</span>*/}
      {/*</HorizontalGroup>*/}
      <HorizontalGroup>
        <h3>Status Data</h3>{" "}
        <div data-testid={testIds.pageOne.status}>
          <pre>{JSON.stringify(statusData, null, 2)}</pre>
        </div>
      </HorizontalGroup>
      <HorizontalGroup>
        <h3>Templates</h3>{" "}
        <div data-testid={testIds.pageOne.templates}>
          <pre>{JSON.stringify(templates, null, 2)}</pre>
        </div>
      </HorizontalGroup>
    </div>
  );
};

// function renderHealth(message: string | undefined) {
//   switch (message) {
//     case "ok":
//       return <Badge color="green" text="OK" icon="heart" />;
//
//     default:
//       return <Badge color="red" text="BAD" icon="bug" />;
//   }
// }
