export default {
  '/api/v1/obtenants': {
    data: [
      {
        charset: 'string',
        clusterName: 'string',
        createTime: 'string',
        locality: 'string',
        name: 'test',
        namespace: 'test',
        status: 'string',
        tenantName: 'string',
        tenantRole: 'string',
        topology: [
          {
            iopsWeight: 0,
            logDiskSize: 'string',
            maxIops: 0,
            memorySize: 'string',
            cpuCount: 1,
            minIops: 0,
            priority: 0,
            type: 'string',
            zone: 'string',
          },
        ],
        unitNumber: 0,
      },
    ],
    message: 'string',
    successful: true,
  },
  '/api/v1/obtenants/test/test': {
    data: {
      charset: 'string',
      clusterName: 'string',
      createTime: 'string',
      locality: 'string',
      name: 'test',
      namespace: 'test',
      primaryTenant: 'string',
      restoreSource: {
        archiveSource: 'string',
        bakDataSource: 'string',
        bakEncryptionSecret: 'string',
        ossAccessSecret: 'string',
        type: 'string',
        until: 'string',
      },
      rootCredential: 'string',
      standbyROCredentail: 'string',
      status: 'string',
      tenantName: 'string',
      tenantRole: 'Primary',
      topology: [
        {
          iopsWeight: 0,
          logDiskSize: 'string',
<<<<<<< Updated upstream
          maxCPU: 'string',
          maxIops: 0,
          memorySize: 'string',
          minCPU: 'string',
=======
          cpuCount: 1,
          maxIops: 0,
          memorySize: 'string',
>>>>>>> Stashed changes
          minIops: 0,
          priority: 0,
          type: 'string',
          zone: 'string',
        },
      ],
      unitNumber: 0,
    },
    message: 'string',
    successful: true,
  },
<<<<<<< Updated upstream
=======
  '/api/v1/obtenants/test/test/backupPolicy': {
    data: {
      archivePath: 'string',
      bakDataPath: 'string',
      bakEncryptionSecret: 'string',
      destType: 'string',
      jobKeepWindow: 'string',
      name: 'string',
      namespace: 'string',
      ossAccessSecret: 'string',
      pieceInterval: 'string',
      recoveryWindow: 'string',
      scheduleDates: [
        {
          backupType: 'typea',
          day: 1,
        },
        {
          backupType: 'typeb',
          day: 2,
        },
        {
          backupType: 'typec',
          day: 3,
        },
      ],
      scheduleTime: 'string',
      scheduleType: 'string',
      status: 'string',
      tenantName: 'string',
    },
    message: 'string',
    successful: true,
  },
  '/api/v1/obtenants/test/test/backup/FULL/jobs': {
    data: [
      {
        encryptionSecret: 'string',
        endTime: 'string',
        name: 'string',
        path: 'string',
        startTime: 'string',
        status: 'string',
        statusInDatabase: 'string',
        type: 'string',
      },
    ],
    message: 'string',
    successful: true,
  },
>>>>>>> Stashed changes
};
