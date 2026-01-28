#!/bin/bash
echo "Starting Frontend Deployment Preparation..."
cd frontend
npm install
npm run build
echo "Build complete. Ready for S3 sync."