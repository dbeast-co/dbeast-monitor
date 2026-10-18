package component_templates

import (
	dataWarehouse "github.com/dbeast/dbeastmonitor/pkg/plugin/data"
)

func init() {
	const componentName = "dbeast-mon-index-corrupted-data-mapping-component"
	const dbeastMinimalCompatibleVersion = "2.0.0"
	dataWarehouse.AppendComponentTemplate(componentName, dbeastMonIndexCorruptedDataMappingComponentContent, dbeastMinimalCompatibleVersion)
}

const dbeastMonIndexCorruptedDataMappingComponentContent string = `
{
  "template": {
    "settings": {},
    "mappings": {
      "properties": {
        "@timestamp": {
          "type": "date"
        },
        "elasticsearch": {
          "properties": {
            "error": {
              "properties": {
                "message": {
                  "type": "text"
                }
              }
            },
            "host": {
              "properties": {
                "hostname": {
                  "type": "keyword",
                  "ignore_above": 256
                }
              }
            },
            "http": {
              "properties": {
                "request": {
                  "properties": {
                    "method": {
                      "type": "keyword",
                      "ignore_above": 256
                    }
                  }
                }
              }
            },
            "node": {
              "properties": {
                "stats": {
                  "properties": {
                    "thread_pool": {
                      "properties": {
                        "@version": {
                          "type": "keyword",
                          "ignore_above": 256
                        },
                        "error": {
                          "properties": {
                            "message": {
                              "type": "text"
                            }
                          }
                        },
                        "host": {
                          "properties": {
                            "hostname": {
                              "type": "keyword",
                              "ignore_above": 256
                            }
                          }
                        },
                        "http": {
                          "properties": {
                            "request": {
                              "properties": {
                                "method": {
                                  "type": "keyword",
                                  "ignore_above": 256
                                }
                              }
                            }
                          }
                        },
                        "tags": {
                          "type": "text",
                          "fields": {
                            "keyword": {
                              "type": "keyword",
                              "ignore_above": 256
                            }
                          }
                        },
                        "url": {
                          "properties": {
                            "full": {
                              "type": "text",
                              "fields": {
                                "keyword": {
                                  "type": "keyword",
                                  "ignore_above": 256
                                }
                              }
                            }
                          }
                        }
                      }
                    }
                  }
                }
              }
            },
            "tags": {
              "type": "keyword",
              "ignore_above": 256
            },
            "url": {
              "properties": {
                "full": {
                  "type": "text",
                  "fields": {
                    "keyword": {
                      "type": "keyword",
                      "ignore_above": 256
                    }
                  }
                }
              }
            }
          }
        },
        "event": {
          "properties": {
            "duration": {
              "type": "long"
            }
          }
        }
      }
    }
  },
  "_meta": {
    "description": "The default mappings for the corrupted data",
    "version": "%DBEAST_VERSION%"
  }
}
`
