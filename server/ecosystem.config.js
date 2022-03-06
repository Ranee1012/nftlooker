module.exports = {
  apps: [
    {
      name: 'nftlooker',
      script: './src/server.js',
      instances: '1',
      env: {
        NODE_ENV: 'development',
      },
      env_production: {
        NODE_ENV: 'production',
      },
    },
  ],
};
