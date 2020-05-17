const isProd = process.env.NODE_ENV === "production";

module.exports = {
    webpack: (config, { isServer }) => {
      // Fixes npm packages that depend on `fs` module
      if (!isServer) {
        config.node = {
          fs: 'empty'
        }
      }
  
      return config
    },
    env: {
      API_ENDPOINT: isProd? "https://3ddcf818.ngrok.io" : "http://localhost:8080"
    }
  }