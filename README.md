# Distributed Tracing Pattern with Istio

Sample applications for learning Distributed Tracing to use Istio.

## Overview

The application is consisting four microservices.

* customers service
  * Return information about the account such as name, email, address, current review and bought item
  * Call histories service and reviews service
* histories service
  * Return information about the history of account watched movie
  * Call items service
* reviews service
  * Return information about the review of movie
  * Call items service
* items service
  * Return information about the item
