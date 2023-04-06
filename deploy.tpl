apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: harmless
  name: harmless
  namespace: harmhaemlesslessNAMESPACE
spec:
  replicas: REPLICANUM
  selector:
    matchLabels:
      app: harmless
  template:
    metadata:
      labels:
        app: harmless
    spec:
      containers:
        - image: 'harmhaemlessIMAGENAME:<BUILD_TAG>'
          imagePullPolicy: IfNotPresent
          name: harmless
          ports:
            - containerPort: 8080
              protocol: TCP
          resources:
            limits:
              memory: 4096Mi
            requests:
              memory: 1024Mi
          env:
            - name: TZ
              value: Asia/Shanghai
            - name: APPENV
              value: harmhaemlessAPPENV           
      imagePullSecrets:
        - name: my-secret
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      terminationGracePeriodSeconds: 30
---
apiVersion: v1
kind: Service
metadata:
  name: label-sys-svc
  namespace: harmhaemlessNAMESPACE
spec:
  ports:
    - name: http
      port: 8080
      protocol: TCP
      targetPort: 8080
  selector:
    app: harmless
  sessionAffinity: None
  type: ClusterIP
status:
  loadBalancer: {}
