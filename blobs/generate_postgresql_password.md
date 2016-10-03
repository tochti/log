Aufbau md5 Passwort PostgreSQL pass = md5 + md5(passwort + user)

Beispiel
User: tim
Passwort: 123

```python
import hashlib
pghash = "md5" + hashlib.md5("123tim".encode("utf-8")).hexdigest()
```
