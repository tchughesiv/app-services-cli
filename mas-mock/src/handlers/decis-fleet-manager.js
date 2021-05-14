module.exports = {
  createServiceAccount: async (c, req, res) => {
    const clientId = Number.MAX_SAFE_INTEGER - new Date().getTime();
    const clientSecret = Number.MAX_SAFE_INTEGER - new Date().getTime();
    res.status(200).json({
      name: req.body.name,
      description: req.body.description,
      clientID: clientId.toString(),
      clientSecret: clientSecret.toString(),
    });
  },
  createDecision: async (c, req, res) => {
    res.status(202).json({
      id: "1iSY6RQ3JKI8Q0OTmjQFd3ocFRg",
      kind: "decision",
      href: "/api/daas-api/v1/decisions/1iSY6RQ3JKI8Q0OTmjQFd3ocFRg",
      status: "complete",
      cloud_provider: "aws",
      multi_az: false,
      region: "us-east-1",
      owner: "api_decision_service",
      name: "serviceapi",
      bootstrapServerHost:
        "serviceapi-1isy6rq3jki8q0otmjqfd3ocfrg.apps.ms-bttg0jn170hp.x5u8.s1.devshift.org",
      created_at: "2020-10-05T12:51:24.053142Z",
      updated_at: "2020-10-05T12:56:36.362208Z",
    });
  },

  deleteDecisionById: async (c, req, res) => {
    res.status(204).json({
      id: "1iSY6RQ3JKI8Q0OTmjQFd3ocFRg",
      kind: "decision",
      href: "/api/daas-api/v1/decisions/1iSY6RQ3JKI8Q0OTmjQFd3ocFRg",
      status: "complete",
      cloud_provider: "aws",
      multi_az: false,
      region: "us-east-1",
      owner: "api_decision_service",
      name: "serviceapi",
      bootstrapServerHost:
        "serviceapi-1isy6rq3jki8q0otmjqfd3ocfrg.apps.ms-bttg0jn170hp.x5u8.s1.devshift.org",
      created_at: "2020-10-05T12:51:24.053142Z",
      updated_at: "2020-10-05T12:56:36.362208Z",
    });
  },

  getDecisionById: async (c, req, res) => {
    res.status(200).json({
      "kind": "Decision",
      "id": "f5e155b9-e1f4-4768-a99e-fc04c5036f0a",
      "version": "4",
      "href": "/api/daas-api/v1/decisions/f5e155b9-e1f4-4768-a99e-fc04c5036f0a/versions/4",
      "name": "mock-decision",
      "description": "A human readable description of my decision",
      "model": {
        "md5": "8eb41527a8f53e5d673771fa2159edac",
        "href": "/api/daas-api/v1/decisions/f5e155b9-e1f4-4768-a99e-fc04c5036f0a/versions/4/dmn"
      },
      "configuration": {
        "key": "value"
      },
      "tags": {
        "key": "value"
      },
      "status": "FAILED",
      "status_message": "Failed to deploy Decision.",
      "submitted_at": "2020-10-05T12:51:24.053142Z",
    });
  },

  listDecisions: async (c, req, res) => {
    return res.status(200).json({
      "kind": "DecisionList",
      "page": 1,
      "size": 2,
      "total": 2,
      "items": [
        {
          "kind": "Decision",
          "id": "f5e155b9-e1f4-4768-a99e-fc04c5036f0a",
          "version": "4",
          "href": "/api/daas-api/v1/decisions/f5e155b9-e1f4-4768-a99e-fc04c5036f0a/versions/4",
          "name": "mock-decision",
          "description": "A human readable description of my decision",
          "model": {
            "md5": "8eb41527a8f53e5d673771fa2159edac",
            "href": "/api/daas-api/v1/decisions/f5e155b9-e1f4-4768-a99e-fc04c5036f0a/versions/4/dmn"
          },
          "configuration": {
            "key": "value"
          },
          "tags": {
            "key": "value"
          },
          "status": "FAILED",
          "status_message": "Failed to deploy Decision.",
          "submitted_at": "2020-10-05T12:51:24.053142Z"
        },
        {
          "kind": "Decision",
          "id": "aba2a420-2afe-456c-b126-073bff1a2023",
          "version": "1",
          "href": "/api/daas-api/v1/decisions/aba2a420-2afe-456c-b126-073bff1a2023/versions/1",
          "name": "mock-decision2",
          "description": "A human readable description of my decision",
          "model": {
            "md5": "8eb41527a8f53e5d673771fa2159edac",
            "href": "/api/daas-api/v1/decisions/aba2a420-2afe-456c-b126-073bff1a2023/versions/1/dmn"
          },
          "configuration": {
            "key": "value"
          },
          "tags": {
            "key": "value"
          },
          "status": "BUILDING",
          "submitted_at": "2020-10-05T12:52:44.053142Z"
        },
      ]
    })
  },

  listCloudProviders: async (_c, _req, res) => {
    return res.status(200).json({
      "kind": "CloudProviderList",
      "page": 1,
      "size": 7,
      "total": 7,
      "items": [
        {
          "kind": "CloudProvider",
          "id": "aws",
          "display_name": "Amazon Web Services",
          "name": "aws",
          "enabled": true
        },
        {
          "kind": "CloudProvider",
          "id": "azure",
          "display_name": "Microsoft Azure",
          "name": "azure",
          "enabled": false
        },
      ]
    })
  },

  listCloudProviderRegions: async (_c, _req, res) => {
    return res.status(200).json(
      {
        "kind": "CloudRegionList",
        "page": 1,
        "size": 17,
        "total": 17,
        "items": [
          {
            "kind": "CloudRegion",
            "id": "eu-west-2",
            "display_name": "EU, London",
            "enabled": true
          }
        ]
      })
  },

  // Handling auth
  notFound: async (c, req, res) => res.status(404).json({ err: "not found" }),
  unauthorizedHandler: async (c, req, res) =>
    res.status(401).json({ err: "unauthorized" }),
};
