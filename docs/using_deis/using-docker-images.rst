:title: Deploying Docker Images on Deis
:description: How to deploy applications on Deis using Docker Images

.. _using-docker-images:

Using Docker Images
===================
Deis supports deploying applications via an existing `Docker Image`_.
This is useful for integrating Deis into Docker-based CI/CD pipelines.

Prepare an Application
----------------------
Start by cloning an example application:

.. code-block:: console

    $ git clone https://github.com/deis/example-go.git
    $ cd example-go
    $ git checkout docker

Next use your local ``docker`` client to build the image and push
it to `DockerHub`_.

.. code-block:: console

    $ docker build -t <username>/example-go .
    $ docker push <username>/example-go

Docker Image Requirements
^^^^^^^^^^^^^^^^^^^^^^^^^
In order to deploy Docker images, they must conform to the following requirements:

 * The Docker image must EXPOSE only one port
 * The exposed port must be an HTTP service that can be connected to an HTTP router
 * A default CMD must be specified for running the container

.. note::

    Docker images which expose more than one port will hit `issue 1156`_.

.. attention::

    Support for non-HTTP services is coming soon

Create an Application
---------------------
Use ``deis create`` to create an application on the :ref:`controller`.

.. code-block:: console

    $ mkdir -p /tmp/example-go && cd /tmp/example-go
    $ deis create
    Creating application... done, created example-go

.. note::

    The ``deis`` client uses the name of the current directory as the
    default app name.

Deploy the Application
----------------------
Use ``deis pull`` to deploy your application from `DockerHub`_ or
a private registry.

.. code-block:: console

    $ deis pull gabrtv/example-go:latest
    Creating build...  done, v2

    $ curl -s http://example-go.local3.deisapp.com
    Powered by Deis

Because you are deploying a Docker image, the ``cmd`` process type is automatically scaled to 1 on first deploy.

.. attention::

    Support for Docker registry authentication is coming soon

Define Process Types
--------------------
Docker containers have a default command usually specified by a `CMD instruction`_.
Deis uses the ``cmd`` process type to refer to this default command.

Process types other than ``cmd`` are not supported when using Docker images.


.. _`Docker Image`: http://docs.docker.io/introduction/understanding-docker/
.. _`DockerHub`: https://registry.hub.docker.com/
.. _`CMD instruction`: http://docs.docker.io/reference/builder/#cmd
.. _`issue 1156`: https://github.com/deis/deis/issues/1156
