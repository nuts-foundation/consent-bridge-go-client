Nuts consent bridge go client
=============================

Library for contacting consent-bridge. Exposed as Nuts-go engine.

.. image:: https://circleci.com/gh/nuts-foundation/consent-bridge-go-client.svg?style=svg
    :target: https://circleci.com/gh/nuts-foundation/consent-bridge-go-client
    :alt: Build Status

.. image:: https://codecov.io/gh/nuts-foundation/consent-bridge-go-client/branch/master/graph/badge.svg
    :target: https://codecov.io/gh/nuts-foundation/consent-bridge-go-client

.. image:: https://api.codeclimate.com/v1/badges/72a11cae5531100dbbbb/maintainability
   :target: https://codeclimate.com/github/nuts-foundation/consent-bridge-go-client/maintainability
   :alt: Maintainability

.. include:: docs/pages/development/consent-bridge-go-client.rst
    :start-after: .. marker-for-readme

Configuration
*************

The following configuration parameters are available:

.. include:: README_options.rst

As with all other properties for nuts-go, they can be set through yaml:

.. sourcecode:: yaml

    cbridge:
       address: localhost:1323

as commandline property

.. sourcecode:: shell

    ./nuts --cbridge.address localhost:1323

Or by using environment variables

.. sourcecode:: shell

    NUTS_CBRIDGE_ADDRESS=localhost:1323 ./nuts
