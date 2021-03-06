package dialogflow

var input1 = []byte(`{
"originalRequest":{
  "source":"google",
  "data":{
    "surface":{
      "capabilities":[
        {"name":"actions.capability.AUDIO_OUTPUT"}
      ]
    },
    "inputs":[{
      "arguments":[],
      "intent":"assistant.intent.action.MAIN",
      "raw_inputs":[{
        "query":"talk to Mokera",
        "input_type":2,
        "annotation_sets":[]
      }]
    }],
    "user":{
      "accessToken":"a75c98e4dbc6f53856ab527f24a62237e7",
      "userId":"1502731363742",
      "permissions":[],
      "locale":"en-US"
    },
    "device":{},
    "is_in_sandbox":true,
    "conversation":{
      "conversation_id":"1502731363742",
      "type":1
    }
  }
},
"id":"ff0f8352-772a-4af9-b1d2-db7d2c1ca8e0",
"timestamp":"2017-08-14T17:22:43.865Z",
"lang":"en",
"result":{
  "source":"agent",
  "resolvedQuery":"GOOGLE_ASSISTANT_WELCOME",
  "speech":"",
  "action":"",
  "actionIncomplete":false,
  "parameters":{},
  "contexts":[
    {"name":"google_assistant_welcome","parameters":{},"lifespan":0},
    {"name":"actions_capability_audio_output","parameters":{},"lifespan":0},
    {"name":"google_assistant_input_type_voice","parameters":{},"lifespan":0}
  ],
  "metadata":{
    "intentId":"3f8d8e18-dd37-4128-a14c-37c52ce4f560",
    "webhookUsed":"true",
    "webhookForSlotFillingUsed":"false",
    "nluResponseTime":1,
    "intentName":"Default Welcome Intent"
  },
  "fulfillment":{
    "speech":"Hi!",
    "messages":[
      {"type":0,"speech":"Hi!"}
    ]
  },
  "score":1.0
},
"status":{
  "code":200,
  "errorType":"success"
},
"sessionId":"1502731363742"}`)

var input2 = []byte(`{
"originalRequest":{
  "source":"google",
  "data":{
    "surface":{
      "capabilities":[{
        "name":"actions.capability.AUDIO_OUTPUT"
      }]
    },
    "inputs":[{
      "arguments":[{
        "raw_text":"Add milk on today",
        "text_value":"Add milk on today",
        "name":"text"
      }],
      "intent":"assistant.intent.action.TEXT",
      "raw_inputs":[{
        "query":"Add milk on today",
        "input_type":2,
        "annotation_sets":[]
      }]
    }],
    "user":{
      "userId":"1502774951934",
      "accessToken":"a75c98e43aba71dbc6f53b527f24a62237e7",
      "permissions":[],
      "locale":"en-US"
    },
    "device":{},
    "is_in_sandbox":true,
    "conversation":{
      "conversation_token":"[]",
      "conversation_id":"1502774951934",
      "type":2
    }
  }
},
"id":"7413ee2c-c3ae-4941-9528-86cfdc926ed9",
"timestamp":"2017-08-15T05:29:35.058Z",
"lang":"en",
"result":{
  "source":"agent",
  "resolvedQuery":"Add milk on today",
  "speech":"",
  "action":"",
  "actionIncomplete":false,
  "parameters":{
    "date":"2017-08-15",
    "event":"milk"
  },
  "contexts":[{
    "name":"google_assistant_input_type_voice",
    "parameters":{
      "date":"2017-08-15",
      "event.original":"milk",
      "date.original":"today",
      "event":"milk"
    },
    "lifespan":0
  },{
    "name":"actions_capability_audio_output",
    "parameters":{
      "date":"2017-08-15",
      "event.original":"milk",
      "date.original":"today",
      "event":"milk"
    },
    "lifespan":0
  }],
  "metadata":{
    "intentId":"9457aa37-9f2a-41ce-8239-ad0d4a94f050",
    "webhookUsed":"true",
    "webhookForSlotFillingUsed":"true",
    "nluResponseTime":98,
    "intentName":"Add event"
  },
  "fulfillment":{
    "speech":"",
    "messages":[{
      "type":0,
      "speech":""
    }]
  },
  "score":1.0
},
"status":{
  "code":200,
  "errorType":"success"
},
"sessionId":"1502774951934"
}`)
