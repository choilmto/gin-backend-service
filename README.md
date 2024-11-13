# README    [![Website shields.io](https://img.shields.io/website-up-down-green-red/http/shields.io.svg)](https://gin-backend-service.onrender.com/)

## About

This backend service is written in golang and uses the gin framework. The service can be used to save and retrieve information about vintage records.

## CI

See 'Actions' tab

## Local development setup

Needs Go 1.18 or later

## API


    /albums

        GET – Get a list of all albums, returned as JSON.
        POST – Add a new album from request data sent as JSON.

    /albums/:id

        GET – Get an album by its ID, returning the album data as JSON.
